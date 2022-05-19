package main

import "fmt"

func main() {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	//send
	go send(even, odd, quit)

	//receive
	receive(even, odd, quit)

	fmt.Println("Above to exit")
}

func send(_even, _odd, _quit chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			_even <- i
		} else {
			_odd <- i
		}
	}
	_quit <- 0
}

func receive(_even, _odd, _quit <-chan int) {
	for {
		select {
		case v := <-_even:
			fmt.Println("From the eve channel:", v)
		case v := <-_odd:
			fmt.Println("from the odd channel:", v)
		case v := <-_quit:
			fmt.Println("from the quit channel:", v)
			return
		}
	}
}
