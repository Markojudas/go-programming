package main

import "fmt"

//struct --> new type with underlying type struct
//kinda like class and objects but in go is a type and value of said type
type person struct {
	first string
	last  string
	age   int
}

//secret agent struct that is a person with a license to kill
type secretAgent struct {
	person //no field name, just the type
	//last   string //to create collision
	ltk bool
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

	fmt.Println()
	fmt.Println("Creating a value of Secret Agent")

	//using the person p1 already created
	sa1 := secretAgent{
		person: p1,
		ltk:    true,
	}

	fmt.Println(sa1)
	fmt.Println(sa1.first, sa1.last, sa1.age, sa1.ltk) //inherited first and last

	fmt.Println()
	fmt.Println("Creating another Secret Agent value; this time with colliding attributes")

	sa2 := secretAgent{
		person: person{
			first: "Jose",
			last:  "Hernandez",
			age:   37,
		},
		//last: "Moquete",
		ltk: true,
	}

	fmt.Println(sa2)
	fmt.Println(sa2.first /*sa2.person.last,*/, sa2.last, sa2.age, sa2.ltk) //inherited first and last; see how we're accessing the 'last' attribute

}
