package main

import "fmt"

// Callback: passing a func as an argument
// functional programming not recommended in Go.

func main() {

	i := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	s := sum(i...) //gotta unbox

	fmt.Println("all numbers", s)

	//Calling Even
	fmt.Println()
	fmt.Println("Calling even()")
	x := even(sum, i...) //sum has the same signature as the requirement f even

	fmt.Println("Even numbers", x)

	//Calling odd
	fmt.Println()
	fmt.Println("Calling odd()")
	y := odd(sum, i...)
	fmt.Println("Odd Numbers:", y)

}

func sum(x ...int) int {

	fmt.Printf("%T\n", x)

	total := 0

	for _, val := range x {
		total += val
	}

	return total
}

func even(f func(x ...int) int, y ...int) int {

	//the function even takes the params : f as a function that takes a slice of int and another slice of int
	// f returns a int
	//the function returns an int as well
	var yi []int
	for _, value := range y {
		if value%2 == 0 {
			yi = append(yi, value)
		}
	}

	return f(yi...) //return an int
}

func odd(f func(x ...int) int, y ...int) int {

	var temp []int

	for _, val := range y {
		if val%2 != 0 {
			temp = append(temp, val)
		}
	}

	return f(temp...)
}
