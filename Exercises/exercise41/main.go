package main

import "fmt"

/*
	Closure is wen we have "enclosed" the scope of a variable in some code block. For this exercise,
		* create a func which "encloses" the scope of a variable
*/

func main() {

	f := foo()
	g := foo()

	fmt.Println(g())
	fmt.Println(f())
	fmt.Println(g())
	fmt.Println(g())
	fmt.Println(f())

}

func foo() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}
