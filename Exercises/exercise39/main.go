package main

import "fmt"

/*
	Create a fucn which returns a func
	Assign the returned func to a variable
	call the return func
*/

func main() {

	x := foo()

	x()

}

func foo() func() {

	return func() {
		fmt.Println("Born into ashes to lose all the games, with a smiling face")
	}
}
