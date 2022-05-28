package main

import "fmt"

/*
	Create a struct "customErr" which implements the builtin.error interface. Create a func "foo" that has a value of type
	error as param. Create a value of type "customErr" and pass it into "foo"

*/

type customErr struct {
	info string
}

//to implement builtin.error
func (ce customErr) Error() string {
	return fmt.Sprintf("Here's the error: %v", ce.info)
}

func main() {

	c1 := customErr{
		info: "I should get more coffee",
	}

	foo(c1)

}

func foo(e error) {
	fmt.Println("foo ran -", e)
	//fmt.Println("foo ran -", e.Error()) -> same output
	//fmt.Println("foo ran -", e.(customErr).info) -> asertion! not the same as conversion
}
