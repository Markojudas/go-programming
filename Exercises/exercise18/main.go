package main

import "fmt"

// Create a program that uses a switch statement with no switch expression specified

func main() {
	x := 42

	switch {
	case x != 42:
		fmt.Println("Do you even know the meaning of life?")
	case x > 42:
		fmt.Println("Too much meaning of life, dial it down!")
	case x < 42:
		fmt.Println("Just doing this here. Supposedly the default can be anywhere.. nope, first or last")
	default:
		fmt.Println("Just let it be. 42 is the meaning")
	}
}
