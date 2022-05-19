package main

import "fmt"

func main() {

	c := make(chan int) //bidirectional channel

	go func() {
		for i := 0; i <= 5; i++ {
			c <- i
		}
		close(c) // if we don't close it will be a deadlock because we are still waiting for more input
	}()

	//looping through the channel until closed
	for v := range c {
		fmt.Println(v)
	}

	fmt.Println("about to exit")
}
