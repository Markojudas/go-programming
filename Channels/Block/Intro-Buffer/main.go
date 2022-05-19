package main

/*
	Buffered Channel
	Coordinate with buffers!
*/

import "fmt"

func main() {
	c := make(chan int, 1) //channels were we can put in ints; this is buffer channel

	c <- 42 // 42 got put in the channel (send)

	fmt.Println(<-c) //(receive)
}
