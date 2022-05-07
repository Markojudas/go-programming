package main

import "fmt"

var y = 42 // meaning of life :)
var z = "Shaken, not stirred"
var a string = `James said, "Shaken, not stirred"` //string literal; keeps spaces and new lines!

var sl string = `Here I am
to prove you
that this is 3 lines`

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y) //prints the type

	fmt.Println(z)
	fmt.Printf("%T\n", z) //prints the type

	fmt.Println(a)
	fmt.Printf("%T\n", a)

	fmt.Println(sl)
	fmt.Printf("%T\n", sl)

	// z = 43 this is a no-go; Go is statically typed.
	// z only holds type string
	// the variable was DECLARED with the IDENTIFIER z
	// and of TYPE string.
}
