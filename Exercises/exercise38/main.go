package main

import "fmt"

// Assign a func to a variable, then call that func

func main() {

	x := func() {
		fmt.Println("I am not a King I am just a bard")
	}
	x()
}
