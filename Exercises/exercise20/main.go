package main

import "fmt"

/*
	Using COMPOSITE LITERAL:
	* create an ARRAY which holds 5 values of TYPE int
	* Assign Values to each index position

	* range over the array and print the values out
	* print the TYPE of the array
*/

func main() {

	arr := [5]int{42, 43, 45, 46, 69} //created and assigned values

	//range over and print
	for i := range arr {
		fmt.Println(arr[i])
	}

	//TYPE
	fmt.Printf("Type of the array %T\n", arr)

}
