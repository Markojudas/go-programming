package main

import "fmt"

/*

	Read Blog:

	https://go.dev/blog/defer-panic-and-recover


	Recover is a built-in function that regains control of a panicking goroutine. Recover is only useful inside deferred functions. During normal execution, a call to recover will return nil and have no other effect. If the current goroutine is panicking, a call to recover will capture the value given to panic and resume normal execution.

*/

func main() {
	f()
	fmt.Println("Returned normally from f")
}

func f() {
	defer func() {
		//Recover only works in deferred functions
		//here we recover the panic value (set in g())
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	fmt.Println("Calling g.")
	g(0)
	//this following statement will not run since we panicked
	fmt.Println("Returned normally from g")
}

func g(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		//storing the panic value as a string
		panic(fmt.Sprintf("%v", i)) //4
	}

	defer fmt.Println("Defer in g", i)
	fmt.Println("Printing g", i)
	g(i + 1)
}
