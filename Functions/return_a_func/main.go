package main

import "fmt"

// Return a function from a function

func main() {
	s1 := foo()
	fmt.Println(s1)

	// assigned the return func (func() int) to the variable X
	x := bar()

	//x is of TYPE func() int
	fmt.Printf("%T\n", x)

	//Using the returned func (assigned to X) and assigning the return of the returned func to i (45)
	//executing the returned func
	i := x()
	fmt.Println(i)

	//we can do execute x() without assigning it to a variable
	fmt.Println(x())

	//or we can even do this
	fmt.Println(bar()())

}

func foo() string {
	s := "hello world"
	return s
}

//returns a function that returns an int

//keyword function name(params) return(s)//this is example it is func() int since it is returning a function that returns an int
func bar() func() int {
	//returns an anonymous function
	return func() int {
		return 451
	}
}
