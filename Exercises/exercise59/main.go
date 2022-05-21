package main

import "fmt"

/*
	write a program that:
		* puts 100 numbers to a channel
		* pulls the numbers off the channel and print them
*/

func main() {
	ch := make(chan int)

	go sender(ch)

	receiver(ch)

	fmt.Println("DONE! Exiting program")

}

func sender(ch chan<- int) {

	for i := 1; i <= 100; i++ {
		ch <- i
	}
	close(ch)
}

func receiver(ch <-chan int) {

	for val := range ch {

		fmt.Println(val)
	}
}
