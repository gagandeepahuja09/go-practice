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
	refillRateInSeconds int
}

func NewTokenBucketRateLimiter(capacity, refillRate int) TokenBucketRateLimiter {
	return TokenBucketRateLimiter{
		capacity:            capacity,
		refillRateInSeconds: refillRate,
	}
}

func (limiter *TokenBucketRateLimiter) Allow(userId string) bool {
	script := `
		local tokens = tonumber(redis.call("get", KEYS[1]))
		local lastRefill = tonumber(redis.call("get", KEYS[2]))

		local now = tonumber(ARGV[1])
		local refillRate = tonumber(ARGV[2])
		local capacity = tonumber(AGRV[3])

		if not lastRefill then
			lastRefill = now
			tokens = capacity
		end

		local timeElapsed = now - lastRefill
		local tokensToAdd = math.floor(timeElapsed * refillRate)
		tokens = math.min(tokens + tokensToAdd, capacity)

		redis.call("set", KEYS[1], tokens)
		redis.call("set", KEYS[2], lastRefill)

		return tokens > 0
	`

	userIdTokensKey := fmt.Sprintf("user_id.%s.tokens", userId)
	userIdLastRefillKey := fmt.Sprintf("user_id.%s.last_refill", userId)

	cmd := redisClient.Eval(context.Background(), script,
		[]string{userIdTokensKey, userIdLastRefillKey},
		time.Now().Unix(), limiter.refillRateInSeconds, limiter.capacity)

	if cmd.Err() != nil {
		return false
	}

	return cmd.Val().(int64) == 1
}

func main() {
	redisClient = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
}
