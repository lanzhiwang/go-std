package main

import (
	"context"
	"fmt"
	"time"
)

func doSomething(ctx context.Context) {
	// func WithCancel(parent Context) (ctx Context, cancel CancelFunc)
    ctx, cancelCtx := context.WithCancel(ctx)

    printCh := make(chan int)
    go doAnother(ctx, printCh)

    for num := 1; num <= 3; num++ {
        printCh <- num
    }

    cancelCtx()

    time.Sleep(100 * time.Millisecond)

    fmt.Printf("doSomething: finished\n")
}

func doAnother(ctx context.Context, printCh <-chan int) {
    for {
        select {
        case <-ctx.Done():
            if err := ctx.Err(); err != nil {
                fmt.Printf("doAnother err: %s\n", err)
            }
            fmt.Printf("doAnother: finished\n")
            return
        case num := <-printCh:
            fmt.Printf("doAnother: %d\n", num)
        }
    }
}

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "myKey", "myValue")
	doSomething(ctx)
}

/*
$ go run main5.go
doAnother: 1
doAnother: 2
doAnother: 3
doAnother err: context canceled
doAnother: finished
doSomething: finished
*/
