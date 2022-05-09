package main

import "fmt"

//built-in types: https://golang.org/
//bool-type : boolean

var x bool

func main() {
	fmt.Println(x) //zero-value : default for bool = false
	x = true
	fmt.Println(x) // true

	// comparisons
	a := 7
	b := 42

	fmt.Println("***Comparisons***")
	fmt.Println(a == b) // false
	fmt.Println(a <= b) // true
	fmt.Println(a > b)  // false
	fmt.Println(a != b) // true

}
