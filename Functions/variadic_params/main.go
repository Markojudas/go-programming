package main

import "fmt"

func main() {
	sum := foo(2, 3, 4, 5, 6, 7, 8, 9)

	fmt.Println("Sum from the main:", sum)
}

func foo(x ...int) int {

	fmt.Println(x)
	fmt.Printf("%T\n", x) //[]int -> slice of ints

	sum := 0

	for _, val := range x {
		sum += val
	}

	return sum
}
