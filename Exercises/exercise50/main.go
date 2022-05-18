package main

import "fmt"

/*
	Method Sets:

	*	Create a type person struct
	*	Attach a method speak to type pointer to a person
	*	Create a type human interface
		* to simplify implement the interface, a human must have the speak method

	*	Create a func "saySomething"
		* have it take a in a human as a param
		* have it call the speak method
	*	Show the following in your code
		* you CAN pass type *person into saySomething
		* you CANNPT pass type person into saySomething
*/

//person struct
type Person struct {
	Name string
	Age  int
}

//interface
type Human interface {
	speak()
}

func (p *Person) speak() {
	name := p.Name

	fmt.Println("Hello, I am,", name)
}

func saySomething(h Human) {
	h.speak()
}

func main() {

	jose := Person{
		Name: "Jose",
		Age:  37,
	}

	fmt.Println("Passing the pointer to", jose.Name, "to the function saySomethin()")
	saySomething(&jose)

	fmt.Println("\nPassing", jose.Name, "to the function this time")
	//saySomething(jose) this is giving an error

	fmt.Println("\nJust calling the speak method")
	jose.speak()
}
