package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func operation1(ctx context.Context) error {
	time.Sleep(1 * time.Second)
	return errors.New("Operation1 failed to run")
}

func operation2(ctx context.Context) {
	// Here we listern for cancellation event.
	select {
	case <-time.After(5 * time.Second):
		fmt.Println("Operation2 completed")
	case <-ctx.Done():
		fmt.Println("Operatrion2 is on HOLD")
	}
}

func main() {
	fmt.Println("Emitting Cancellation event")

	ctx := context.Background()

	ctx, cancle := context.WithCancel(ctx)

	go func() {
		err := operation1(ctx)
		if err != nil {
			cancle()
		}
	}()

	operation2(ctx)
}
