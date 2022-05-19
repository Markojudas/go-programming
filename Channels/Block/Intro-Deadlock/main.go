package main

/*

	Unbuffered Channel
	THIS DOES NOT RUN
	Deadlock - all goroutines are asleep

	Channels Block
*/

import "fmt"

func main() {
	c := make(chan int) //channels were we can put in ints

	c <- 42 //putting 42 to the channel

	fmt.Println(<-c)
}
