package main

import (
	"context"
	"fmt"
	"time"
)

type SlowFunction func(string) (string, error)

type WithContext func(context.Context, string) (string, error)

func Timeout(sf SlowFunction) WithContext {
	return func(ctx context.Context, s string) (string, error) {
		chres := make(chan string)
		cherr := make(chan error)

		go func() {
			res, err := sf(s)
			chres <- res
			cherr <- err
		}()

		select {
		case res := <-chres:
			return res, <-cherr
		case <-ctx.Done():
			return "", ctx.Err()
		}
	}
}

func main() {
	ctx := context.Background()
	ctxt, cancel := context.WithTimeout(ctx, 100*time.Millisecond)
	defer cancel()

	slow := func(s string) (string, error) {
		time.Sleep(200 * time.Millisecond)
		return s, nil
	}

	timeout := Timeout(slow)

	res, err := timeout(ctxt, "some input")

	fmt.Println(res, err)
}
