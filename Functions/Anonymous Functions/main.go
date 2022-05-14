package main

import "fmt"

func main() {

	//regular function
	foo()

	//anonymous function or self executing function
	func() {
		fmt.Println("Anonymous func ran")
	}() //no arguments

	func(x int) {
		fmt.Println("The meaning of life:", x)
	}(42) //takes 1 argument, passing 42

}

func foo() {
	fmt.Println("Foo Ran")
}
