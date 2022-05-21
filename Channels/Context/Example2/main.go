package main

import (
	"context"
	"fmt"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())

	defer cancel() //will run at the end

	for val := range gen(ctx) {
		fmt.Println(val)
		if val == 5 {
			break
		}
	}
}

func gen(ctx context.Context) <-chan int {
	//returns a receive channel

	dst := make(chan int)
	n := 1

	go func() {
		for {
			select {
			case <-ctx.Done():
				return //returning not to leak the goroutine
			case dst <- n:
				n++
			}
		}
	}()

	return dst
}
