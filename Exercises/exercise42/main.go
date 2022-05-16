package main

import "fmt"

// Create a value and assign it to a variable then print the address of that value

func main() {

	x := "String!"

	fmt.Println(x)
	fmt.Println(&x)

}
