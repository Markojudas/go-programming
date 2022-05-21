package main

import "fmt"

/*
	Using this code: https://go.dev/play/p/YHOMV9NYKK

	Show the comma ok idiom
*/

func main() {

	c := make(chan int) //bidirectional channel

	//using a goroutine send an int to the channel
	go func() {
		c <- 42
	}()

	//get something of the channel; if there's something ok will be true
	v, ok := <-c
	fmt.Println(v, ok)

	close(c) // here we close c lest we deadlock

	v, ok = <-c // at this point there is nothing in the channel and therefore ok == false
	fmt.Println(v, ok)
}
