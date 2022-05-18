package main

import "fmt"

func main() {

	//assigning a function to a variable
	//func is a TYPE!
	f := func() {
		fmt.Println("My First func expression")
	}
	f()

	g := func(x int) {
		fmt.Println("The year big brother started watching:", x)
	}
	g(1984)

}
