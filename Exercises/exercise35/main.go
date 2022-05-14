package main

import "fmt"

/*
	Create a user defined struct with
		* identifier "person"
		* the fields
			* first
			* last
			* age
	* Attach a method to type person with
			* the identifier "speak"
			* the method should have the person say their name and age
	* Create a value of type person
	* Call the method from the value of type person
*/

type person struct {
	first string
	last  string
	age   int
}

func main() {

	pJose := person{
		first: "Jose",
		last:  "Hernandez",
		age:   37,
	}

	pJose.speak()

}

func (p person) speak() {

	fmt.Println("Hello my name is", p.first, p.last, "and I am", p.age, "years old!")
}
