package main

import "fmt"

/*
	a "callback" is when we pass a func into a func as an argument. for this exercise,
		* pass a func into a func as an argumnet
*/

func main() {

	f := func(x int) int {
		return x * x
	}

	ans := foo(f, 5)

	fmt.Println(ans) // 25

}

func foo(f func(x int) int, x int) int {

	return f(x)
}
