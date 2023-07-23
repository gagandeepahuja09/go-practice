package main

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

var redisClient *redis.Client

type TokenBucketRateLimiter struct {
	capacity            int
	refillRatePerSecond int
}

func NewTokenBucketRateLimiter(capacity, refillRate int) TokenBucketRateLimiter {
	return TokenBucketRateLimiter{
		capacity:            capacity,
		refillRatePerSecond: refillRate,
	}
}

func (limiter *TokenBucketRateLimiter) Allow(userId string) bool {
	script := `
		local tokens = tonumber(redis.call("get", KEYS[1]))
		local lastRefill = tonumber(redis.call("get", KEYS[2]))

		local now = tonumber(ARGV[1])
		local refillRate = tonumber(ARGV[2])
		local capacity = tonumber(ARGV[3])

		if not lastRefill then
			lastRefill = now
			tokens = capacity
		end

		-- refill logic
		local secondsElapsed = now - lastRefill
		local tokensToAdd = secondsElapsed * refillRate
		tokens = math.min(tokens + tokensToAdd, capacity)
		lastRefill = now

		-- consumption logic
		tokens = math.max(tokens - 1, -1)

		local ttl = 60 -- only for testing

		redis.call("setex", KEYS[1], ttl, tokens)
		redis.call("setex", KEYS[2], ttl, lastRefill)

		return tokens
	`

	userIdTokensKey := fmt.Sprintf("user_id.%s.tokens", userId)
	userIdLastRefillKey := fmt.Sprintf("user_id.%s.last_refill", userId)

	cmd := redisClient.Eval(context.Background(), script,
		[]string{userIdTokensKey, userIdLastRefillKey},
		time.Now().Unix(), limiter.refillRatePerSecond, limiter.capacity)

	if cmd.Err() != nil {
		fmt.Println(cmd.Err().Error())
		return false
	}

	tokenCount := cmd.Val().(int64)

	fmt.Println("tokenCount", tokenCount)

	return tokenCount >= 0
}

func main() {
	redisClient = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	limiter := NewTokenBucketRateLimiter(10, 2)

	userId := time.Now().GoString()
	for i := 1; i <= 40; i++ {
		if limiter.Allow(userId) {
			fmt.Printf("Request %d allowed\n", i)
		} else {
			fmt.Printf("Request %d rejected\n", i)
		}
		time.Sleep(100 * time.Millisecond)
		if i > 30 {
			time.Sleep(500 * time.Millisecond)
		}
	}
}
