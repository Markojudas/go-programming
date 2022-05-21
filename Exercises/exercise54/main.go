package main

/*
	func main(){
		c := make(chan int)

		c <- 42

		fmt.Println(<-c)
	}

	Get this code working:
		*	Using func literal, aka an anonymous self-executing func
		*	A buffered channel

*/

import (
	"fmt"
)

func main() {

	// Using a func literal
	/* c := make(chan int)

	go func() {
		c <- 42
	}()

	fmt.Println(<-c) */

	c := make(chan int, 1)

	c <- 42

	fmt.Println(<-c)
}
