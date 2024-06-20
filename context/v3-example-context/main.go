package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("got the stop channel")
				return
			default:
				fmt.Println("still working")
				time.Sleep(1 * time.Second)
			}
		}
	}()

	time.Sleep(5 * time.Second)
	fmt.Println("main stop the gorutine")
	cancel()
	time.Sleep(5 * time.Second)
}

// $ ./main
// still working
// still working
// still working
// still working
// still working
// main stop the gorutine
// got the stop channel
// $
