package main

import (
	"context"
	"fmt"
	"time"
)

func task(ctx context.Context) {
	i := 1
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Ending the context")
			fmt.Println(ctx.Err())
			// ctx.Err could be either deadline exceeded or context cancelled
			time.Sleep(time.Second)
		default:
			fmt.Println(i)
			time.Sleep(time.Second)
			i++
		}
	}
}

func main() {
	ctx := context.Background()

	cancelCtx, cancel := context.WithCancel(ctx)

	go task(cancelCtx)
	time.Sleep(time.Second * 4)
	cancel()
	time.Sleep(time.Second * 1)
}
