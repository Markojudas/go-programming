package main

import "fmt"

/*
	Use the "defer" keyword to show that a deferred func runs after the func containing it exits.
*/

func main() {

	fmt.Println("Defering the lastFunc() now but it will be printed Last!")
	defer lastFunc()
	foo()
	fmt.Println("\nFoo just ran and we're back to main()")
	bar()
	fmt.Println("\nBack in Main and the last line. The deferred func will follow:")

}

func foo() {
	fmt.Println("Foo is printing")
}

func bar() {
	fmt.Println("Bar is printing")
}

func lastFunc() {
	fmt.Println("this is the deferred func")
}
