package main

import (
	"context"
	"fmt"
	"time"
)

// Реализовать все возможные способы остановки выполнения горутины.
func main() {
	stop := make(chan bool)

	go workWithChan(stop)
	time.Sleep(5 * time.Second)
	stop <- true

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	go workWithCtx(ctx)
	time.Sleep(4 * time.Second)

	timer := time.NewTimer(2 * time.Second)

	go workWithTimer(timer)
	time.Sleep(3 * time.Second)

	fmt.Println("exit")
}

// With Timer variant
func workWithTimer(timer *time.Timer) {
	for {
		select {
		case <-timer.C:
			fmt.Println("stop with timer")
			return
		default:
			fmt.Println("working with timer")
			time.Sleep(1 * time.Second)
		}
	}
}

// With Context variant
func workWithCtx(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("stop with ctx")
			return
		default:
			fmt.Println("working with ctx")
			time.Sleep(1 * time.Second)
		}
	}
}

// Channel variant
func workWithChan(ch <-chan bool) {
	for {
		select {
		case <-ch:
			fmt.Println("stop")
			return
		default:
			fmt.Println("working")
			time.Sleep(1 * time.Second)
		}
	}
}
