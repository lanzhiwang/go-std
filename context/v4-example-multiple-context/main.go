package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go worker(ctx, "node01")
	go worker(ctx, "node02")
	go worker(ctx, "node03")

	time.Sleep(5 * time.Second)
	fmt.Println("mian stop the gorutine")
	cancel()
	time.Sleep(5 * time.Second)
}

func worker(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, "got the stop channel")
			return
		default:
			fmt.Println(name, "still working")
			time.Sleep(1 * time.Second)
		}
	}
}

// $ ./main
// node03 still working
// node01 still working
// node02 still working
// node02 still working
// node03 still working
// node01 still working
// node01 still working
// node02 still working
// node03 still working
// node03 still working
// node01 still working
// node02 still working
// node02 still working
// node03 still working
// node01 still working
// mian stop the gorutine
// node01 got the stop channel
// node02 got the stop channel
// node03 got the stop channel
// $
