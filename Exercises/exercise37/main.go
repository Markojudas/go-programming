package main

import "fmt"

/*
	Build an use an anonymous func
*/

func main() {

	x := 2

	fmt.Println(x)
	func() {
		x++
	}()

	fmt.Println("After calling the anonymous func:", x)

}
