package main

import "fmt"

//struct --> new type with underlying type struct
//kinda like class and objects but in go is a type and value of said type
type person struct {
	first string
	last  string
	age   int
}

func main() {
	//creating values of that type using composite literal
	p1 := person{
		first: "James",
		last:  "Bond",
		age:   32,
	}

	p2 := person{
		first: "Miss",
		last:  "Moenypenny",
		age:   33,
	}

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println()

	fmt.Println("Accessing the 'fields'")
	fmt.Println(p1.first, p1.last, p1.age)
	fmt.Println(p2.first, p2.last, p2.age)

}
