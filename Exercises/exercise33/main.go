package main

import "fmt"

/*
	Create a func with the IDENTIFIER foo that:
		* Takes a variadic parameter of type int
		* pass in a value of type []int into your func (unfurl the []int)
		* returns the sum of all values of type int passed in
	Create a func with the IDENTIFIER bar that:
		* Takes in a parameter of type []int
		* Returns the sum of all values of type int passed in
*/

func main() {

	arr := []int{1, 2, 3, 4, 5}

	totalA := foo(arr...)
	totalB := bar(arr)

	fmt.Println(totalA)
	fmt.Println(totalB)

}

func foo(x ...int) int {

	var total int

	for _, val := range x {
		total += val
	}

	return total
}

func bar(x []int) int {
	var total int

	for _, val := range x {
		total += val
	}

	return total
}
