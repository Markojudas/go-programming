package main

/*
	Unbuffered Channel
	THIS RUNS
*/

import "fmt"

func main() {
	c := make(chan int) //channels were we can put in ints

	go func() {
		c <- 42 //this is another goroutine so the are another "thread" this works (send to a goroutine)
	}()

	fmt.Println(<-c) //this takes it - coordinated (receive)
}
