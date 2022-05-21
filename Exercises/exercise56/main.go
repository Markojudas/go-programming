package main

/*
	Starting with this code: https://go.dev/play/p/sfyu4Is3FG

		*	Pull the values off the channel using a for range loop
*/

import "fmt"

func main() {
	c := gen()

	receive(c)

	fmt.Println("about to exit")
}

func gen() <-chan int {
	c := make(chan int)

	//had to create an anon func lest the program deadlocks
	go func() {

		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c) //closing the channel
	}()

	return c
}

func receive(c <-chan int) {

	for v := range c {
		fmt.Println(v)
	}

}
