package main

import "fmt"

// Recursion

func main() {

	x := fact(5)
	fmt.Println(x)

	y := factLoop(5)
	fmt.Println(y)

	f := fib(10)
	fmt.Println(f)
}

func fact(n int) int {

	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func factLoop(n int) int {
	var total = 1

	for n > 0 {
		total *= n
		n--
	}

	return total
}

func fib(n int) int {

	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	return fib(n-1) + fib(n-2)
}
