package main

import "fmt"

func main() {

	c := make(chan int) //bidirectional channel

	// send
	go foo(c)

	// receive
	bar(c) //this will block until something is sent! once foo() sends and finishes this will continue

	fmt.Println("about to exit")
}

//send
func foo(c chan<- int) {
	//since we can go from general to specific
	//so we send the channel and treat it as a send-only
	c <- 42
}

//receive
func bar(c <-chan int) {
	//since we can go from general to specific
	//so we send the channel and treat it as a receiver-only
	fmt.Println(<-c)
}
