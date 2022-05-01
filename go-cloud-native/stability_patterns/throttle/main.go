package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Effector func(context.Context) (string, error)

func Throttle(effector Effector, maxTokens uint, refillCount uint, d time.Duration) Effector {
	tokenCount := maxTokens
	var once sync.Once
	var ticker *time.Ticker
	return func(ctx context.Context) (string, error) {
		if ctx.Err() != nil {
			return "", ctx.Err()
		}
		once.Do(func() {
			ticker = time.NewTicker(d)
			for {
				select {
				case <-ctx.Done():
					return
				case <-ticker.C:
					totalTokens := tokenCount + refillCount
					if totalTokens > maxTokens {
						totalTokens = maxTokens
					}
					tokenCount = totalTokens
				}
			}
		})

		if tokenCount <= 0 {
			return "", fmt.Errorf("too many calls")
		}

		tokenCount--

		return effector(ctx)
	}
}

func main() {

}
