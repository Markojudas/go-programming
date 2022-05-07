package main

import "fmt"

var y = 43

// when you need to declare a variable OUTSIDE of a func body use VAR
// Best practice; limit the scope of variables and use short declarations  as much as you can
// the scope of y is package level (global)
// declare & assign = initilization ; dynamic type init

var z int

// DECLARE there is a VARIABLE with the IDENTIFIER 'z'
// and that the VARAIBLE with the IDENTIFIER 'z' is of TYPE int
// ASSIGNS the ZERO (0) VALUE of TYPE int to 'z'
// Each element of such a variable or value is set to the zero value
// for its type: false for booleans, 0 for numeric types, "" for strings,
// and nil for pointers, functions, interfaces, slices, channels, and maps

func main() {
	//short declaration operator
	//DECLARE a variable and ASSIGN a VALUE

	x := 42
	fmt.Println(x)

	fmt.Println(y)

	foo()

	fmt.Println(z)
}

func foo() {
	fmt.Println("Again", y)
}
