package main

/*
	Directional Channels - RECEIVE-ONLY Channel
*/

import "fmt"

func main() {
	c := make(<-chan int, 2) // Receive Only Channel

	//c <- 42 Doesn't run
	//c <- 43 Doesn't run

	fmt.Println(<-c)
	fmt.Println(<-c)

	fmt.Println("-----------")
	fmt.Printf("%T\n", c) //TYPE: chan int
}
