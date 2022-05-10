package main

import "fmt"

// Write a program that prints a number in decimal, binary, and hex
func main() {
	x := 42 //short declaration

	fmt.Printf("The Decimal Representation of X: %d\n", x) //42
	fmt.Printf("The Binary Representation of X: %b\n", x)  // 10 1010
	fmt.Printf("The Hex Representation of X: %#x\n", x)    //  0x2a
	fmt.Printf("The Octal Representation of X: %#o\n", x)  // 052 octal (extra :))
}
