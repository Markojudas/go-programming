package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	fmt.Println("Error Check 1:", ctx.Err())
	fmt.Println("num goroutines 1:", runtime.NumGoroutine())

	go func() {
		n := 0

		for {
			// DONE is provided for use in select statements
			select {
			case <-ctx.Done():
				return //ends the function
			default:
				n++
				time.Sleep(time.Millisecond * 200)
				fmt.Println("Working", n)
			}
		}
	}()

	time.Sleep(time.Second * 2)
	fmt.Println("error check 2:", ctx.Err())
	fmt.Println("Num goroutines 2:", runtime.NumGoroutine())

	fmt.Println("About to cancel Context")
	cancel()
	fmt.Println("Canceled context")

	time.Sleep(time.Second * 2)
	fmt.Println("error check 3:", ctx.Err())
	fmt.Println("Num goroutines 3:", runtime.NumGoroutine())

}
