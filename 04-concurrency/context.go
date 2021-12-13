package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	cancelCtx, cancelFunc := context.WithCancel(ctx)
	defer func() {
		time.Sleep(3 * time.Second)
		fmt.Println("main goroutine start invoke cancelFunc")
		cancelFunc()
		time.Sleep(3 * time.Second)
	}()

	go func(ctx context.Context) {
		defer func() {
			fmt.Println("child goroutine had finished")
		}()
		select {
		case <-ctx.Done():
			fmt.Println("child goroutine receive signal with cancel")
		}
	}(cancelCtx)

}
