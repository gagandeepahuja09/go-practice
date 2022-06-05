package main

import (
	"context"
	"time"
)

// Effector is the function that you want to subject to throttling
type Effector func(context.Context) (string, error)

// Throttled wraps an Effector. It accepts the same parameters, plus a UID string.
// It returns the same, plus a bool that's true if call is not Throttled
type Throttled func(context.Context, string) (bool, string, error)

// A bucket tracks the request associate with a UID.
type bucket struct {
	tokens uint
	time   time.Time
}

// Throttle accepts an Effector function and returns a Throttled function with a
// per-UID token bucket with a capacity of max that refills at a rate of refill
// tokens every d duration.
func Throttle(e Effector, max uint, refill uint, d time.Duration) Throttled {
	// bucket maps UIDs to specific buckets
	buckets := map[string]*bucket{}

	return func(ctx context.Context, uid string) (bool, string, error) {
		b := buckets[uid]

		// This is a new entry. It passes. Assumes that capacity >= 1
		if b == nil {
			buckets[uid] = &bucket{tokens: max - 1, time: time.Now()}

			str, err := e(ctx)
			return true, str, err
		}

		// Calculate how many tokens we now have based on the time passed since the
		// last request.
		timePassed := time.Since(b.time)
		numRefills := uint(timePassed / d)
		tokensAdded := numRefills * refill
		currentTokens := b.tokens + tokensAdded

		// We don't have enough tokens. Return false.
		if currentTokens < 1 {
			return false, "", nil
		}

		// If we have refilled our bucket, we can restart the clock
		// Otherwise we figure out when the most recent tokens were added.
		// We also decrement tokens by 1 since this is a successful request.
		if currentTokens > max {
			b.time = time.Now()
			b.tokens = max - 1
		} else {
			deltaTime := time.Duration(numRefills) * d

			b.time = b.time.Add(deltaTime)
			b.tokens = currentTokens - 1
		}

		str, err := e(ctx)
		return true, str, err
	}
}

func main() {

}
