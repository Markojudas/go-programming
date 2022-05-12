package main

import "fmt"

/*
	DELETE from a slice, using APPEND and SLICING

	output: [42,43,44,48,49,50,51] // deleted 45,46,47
*/

func main() {

	arrOrig := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println("Original Slice")
	fmt.Println(arrOrig)
	fmt.Println()

	fmt.Println("After deletion...")
	arrOrig = append(arrOrig[:3], arrOrig[6:]...)
	fmt.Println(arrOrig)

}
