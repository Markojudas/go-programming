package main

import "fmt"

var x int
var y string
var z bool

func main() {
	fmt.Printf("%T\n", x) // 0; zero value
	fmt.Printf("%T\n", y) // ""; empty string aka zero value
	fmt.Printf("%T\n", z) // false; zero value
	fmt.Println("\nPrinting the values")
	fmt.Printf("X = %v\n", x)
	fmt.Printf("Y = %v\n", y)
	fmt.Printf("Z = %v\n", z)

}
