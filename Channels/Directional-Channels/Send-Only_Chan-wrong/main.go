package main

/*
	Directional Channels - SEND-ONLY Channel
*/

import "fmt"

func main() {
	c := make(chan<- int, 2) // channel is only a send value; unable to receive (SEND-ONLY CHANNEL)

	c <- 42
	c <- 43

	//fmt.Println(<-c) Doesn't run
	//fmt.Println(<-c) doesn't run

	fmt.Println("-----------")
	fmt.Printf("%T\n", c) //TYPE: chan int
}
