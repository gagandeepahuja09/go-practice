package main

import (
	"context"
	"errors"
	"sync"
	"time"
)

// this can take whatever form but it must return an error.
type Circuit func(context.Context) (string, error)

func Breaker(circuit Circuit, failureThreshold int) Circuit {
	consecutiveFailures := 0
	lastAttempt := time.Now()
	var m sync.RWMutex

	return func(ctx context.Context) (string, error) {
		m.RLock()
		d := consecutiveFailures - failureThreshold
		if d >= 0 {
			shouldRetryAt := lastAttempt.Add(time.Second * 2 << d)
			if !time.Now().After(shouldRetryAt) {
				m.RUnlock()
				return "", errors.New("service unreachable")
			}
		}

		m.RUnlock()

		response, err := circuit(ctx)

		m.Lock()
		defer m.Unlock()

		lastAttempt = time.Now()

		if err != nil {
			consecutiveFailures++
			return response, err
		}

		consecutiveFailures = 0 // Reset failures count

		return response, nil
	}
}

func main() {

}
