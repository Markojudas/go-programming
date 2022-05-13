package main

import "fmt"

func main() {

	defer foo() //it runs AFTER main() exist; bar will run first, main will finish, then foo runs
	//example: you open a file and line below it we can defer the func to close. better organization
	bar()
}

func foo() {
	fmt.Println("foo")
}

func bar() {
	fmt.Println("bar")
}
