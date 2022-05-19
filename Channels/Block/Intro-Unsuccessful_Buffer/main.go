package main

/*
	Buffered Channel
	Coordinate with buffers! This doesn't run! Deadlock
*/

import "fmt"

func main() {
	c := make(chan int, 1) //channels were we can put in ints; this is buffer channel

	c <- 42 // 42 got put in the channel
	c <- 43 // only one room for one until the value is taken out, buffer

	fmt.Println(<-c)
}
