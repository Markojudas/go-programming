package main

import "fmt"

// Create a func with IDENTIFIER foo that returns an int
// Create a func with IDENTIFIER bar that return an int and a string
// Call both funcs
// print out their results

func main() {

	a := foo()
	b, c := bar()

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

}

func foo() int {
	return 42
}

func bar() (int, string) {
	return 69, "nice"
}
