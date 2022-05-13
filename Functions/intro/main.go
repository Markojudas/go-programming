package main

import "fmt"

/*
	syntax:
		func (r receiver) identifier(parameters) (returns(s)){ ... }
*/

func main() {

	foo()
	bar("Jose")
	s1 := woo("Moneypenny")
	x, y := mouse("Ian", "Flemming")

	fmt.Println(s1)
	fmt.Println(x)
	fmt.Println(y)
}

func foo() {
	fmt.Println("Hello from foo")
}

//everything in Go is PASS BY VALUE
func bar(s string) {
	fmt.Println("hello,", s)
}

//returns a value
func woo(s string) string {
	return fmt.Sprint("Hello from woo, ", s) //returning the string
}

//returning multiple values
func mouse(first string, last string) (string, bool) {
	a := fmt.Sprint(first, " ", last, ", says hello!")
	b := false

	return a, b

}
