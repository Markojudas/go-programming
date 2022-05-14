package main

import "fmt"

//closure: enclose the scope of a variable

//package level scope
var x int

func main() {
	var x1 int
	fmt.Println(x)
	x++
	fmt.Println(x)

	foo()
	fmt.Println(x)

	//limiting the scope of x1 to just main()
	fmt.Println()
	fmt.Println(x1)
	fmt.Println("Calling foo() again")
	foo()
	fmt.Println(x)
	fmt.Println(x1)

	//a block of code in a block of code
	{
		var y int = 33
		fmt.Println(y)
	}

	//y++ cannot be accessed "undefined"

	fmt.Println("\nMore realistic example: incrementor()")
	a := incrementor()
	b := incrementor()

	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(b())
	fmt.Println(b())
}

func foo() {
	fmt.Println("NOW ON FOO()")
	x++
	//x1++ is not accessible
}

func incrementor() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}
