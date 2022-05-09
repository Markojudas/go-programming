package main

import "fmt"

type powerlevel int

var x powerlevel

var y int

func main() {
	fmt.Printf("The Value of X: %v\n", x)
	fmt.Printf("The Value of X: %T\n", x)

	x = 42

	fmt.Printf("The value of X after assignment: %v\n", x)

	y = int(x)

	fmt.Printf("The value of Y : %v\n", y)
	fmt.Printf("The Type of Y : %T\n", y)

}
