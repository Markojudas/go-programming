package main

import "fmt"

// create a program that uses a switch statement with the switch expression specified as a variable of TYPE string with the IDENTIFIER "favSport"

func main() {
	var favSport string = "soccer" // [Basketball | Baseball]

	switch favSport {
	case "Basketball":
		fmt.Println("Go shoot hoops")
	case "Baseball":
		fmt.Println("Hit coca-cola")
	default:
		fmt.Println("GOOOOOOOOAAAALLL")
	}
}
