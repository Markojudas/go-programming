package main

import "fmt"

//golang.org or godoc.org
//packages are kinda to organize the code
/*
	variadic parameters
	the "...<some type>" is how we signify a variadic parameter (0 or more parameters)
	"interface{} is the empty interface, every value is also of type interface{}"
	any is an alias for the empty interface
	"a...interface{}" or "a...any" means "give me as many arguments of any type"
*/

func main() {
	n, err := fmt.Println("Hello, playground", 42, true)
	//fmt.Println returns the number of bytes written and any errors
	// every variable needs to be used!
	//if you don't want to use it, through it away!
	fmt.Println(n)
	fmt.Println(err)

	n2, _ := fmt.Println("I am throwing away the error variable")

	fmt.Println("The number of bytes written above", n2)
}
