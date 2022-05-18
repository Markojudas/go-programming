package main

import (
	"fmt"
	"sync"
)

/*
	In addition to the main goroutine (main()) launch two additional goroutines
		* each additional goroutine should print something out
	use waitgroups to make sure each goroutine finishes before your program exists
*/

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	fmt.Println("THIS IS THE MAIN GOROUTINE")

	go foo()
	go bar()

	wg.Wait()

	fmt.Println("exiting!")
}

func foo() {

	fmt.Println("This is the foo goroutine")
	wg.Done()
}

func bar() {
	fmt.Println("this is the bar goroutine")
	wg.Done()
}
