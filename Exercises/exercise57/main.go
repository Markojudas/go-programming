package main

import "fmt"

/*

	Pull the values off the channell using a select statement

	func main() {
		q := make(chan int)
		c := gen(q)

		receive(c, q)

		fmt.Println("about to exit")
	}

	func gen(q <-chan int) <-chan int {
		c := make(chan int)

		for i := 0; i < 100; i++ {
			c <- i
		}

		return c
	}

*/

func main() {

	q := make(chan int) //this is quit
	c := gen(q)

	receiver(c, q)

	fmt.Println("Exiting")
}

func gen(q chan<- int) <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		q <- 1
		close(c)
	}()

	return c
}

func receiver(c, q <-chan int) {

	for {
		select {
		case v := <-c:
			fmt.Println(v)
		case <-q:
			return
		}
	}
}
