package main

import "fmt"

// Array is the building block

func main() {

	// Create
	var x [5]int // Gotta specify the size; same type; length is PART of the TYPE

	fmt.Println(x) // all indexes are initializes with 0

	//assign a value to an index
	x[3] = 42

	fmt.Println(x)

	// get the length
	fmt.Println(len(x))

	//Arrays are building blocks for Slices; use Slices instead

}
