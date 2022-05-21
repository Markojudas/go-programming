package main

import (
	"fmt"
	"runtime"
)

/*
	Write a program that:
		* launches 10 goroutines
			* each goroutine adds 10 numbers to a channel
		* pull the numbers off the channel and print them
*/

func main() {

	ch := make(chan int)

	fmt.Println("Num of goroutines:\t", runtime.NumGoroutine())

	for i := 0; i < 10; i++ {
		go sender(ch)
	}

	receiver(ch)

	fmt.Println("Num of goroutines:\t", runtime.NumGoroutine())
	fmt.Println("About to exit")
}

func sender(ch chan<- int) {

	for i := 0; i < 10; i++ {
		ch <- i
	}
}

func receiver(ch <-chan int) {

	for i := 0; i < 100; i++ {
		fmt.Println(i, ":", <-ch)
	}

}
