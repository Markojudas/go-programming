package main

import "fmt"

/*
	using COMPOSITE LITERAL

	*Create a slice of size 10, assign
	* range over it and print values
	* print type of slice

*/

func main() {
	arr := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	for _, val := range arr {
		fmt.Println(val)
	}

	fmt.Printf("Type of the Slice: %T\n", arr)
}
