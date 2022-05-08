package main

import "fmt"

type powerlevel int

var x powerlevel

func main() {
	fmt.Printf("The Value of X: %v\n", x)
	fmt.Printf("The type of X: %T\n", x)
	x = 42
	fmt.Printf("The Value of X after assignment: %v\n", x)

}
