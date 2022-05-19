package main

/*
	Buffered Channel
	Coordinate with buffers! This doesn't run! Deadlock
*/

import "fmt"

func main() {
	c := make(chan int, 2) //channels were we can put in ints; this is buffer channel

	c <- 42 // 42 got put in the channel (send)
	c <- 43 // since the buffer is size two, this will work

	fmt.Println(<-c) //(receive)
	fmt.Println(<-c)
}
