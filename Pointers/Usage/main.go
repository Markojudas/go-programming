package main

import "fmt"

func main() {

	x := 0
	foo(x)
	fmt.Println(x) // x is still 0; pass by value

	fmt.Println("Passing the address of x")
	fmt.Println("Address to x, before:", &x)
	fmt.Println("Value of x, before:", x)
	bar(&x)
	fmt.Println("Address to x, after:", &x) // same address
	fmt.Println("Value of x, after:", x)
	fmt.Println(x) // now x is 43

}

func foo(y int) {
	fmt.Println(y)
	y = 43
	fmt.Println(y)
}

func bar(y *int) {
	fmt.Println(*y)
	fmt.Println("Address passed:", y)
	*y = 43
	fmt.Println(*y)
}
