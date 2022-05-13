package main

import "fmt"

//interface - "If you have these methods you are also my TYPE". Allows us to take advantage of polymorphism
//empty interface == any value of any type. Empty interface means NO METHODS == any value.

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

// func (r receiver) identifier(params) (return(s)) { code }

//any value of the receiver type -> method is the behavior of an object/TYPE
func (s secretAgent) speak() {
	fmt.Println("I am", s.first, s.last, " - the secret agent speak")
}

func (p person) speak() {
	fmt.Println("I am", p.first, p.last, " - the person speak")
}

//any TYPE that has the method speak() is also of TYPE human
type human interface {
	speak()
}

func bar(h human) {
	//assertion
	switch h := h.(type) {
	case person:
		fmt.Println("I was passed into barrrrrrrrr", h.first)
	case secretAgent:
		fmt.Println("I was passedin barrrr", h.first)
	}
	fmt.Println("I was passed into bar", h)
}

type hotdog int

func main() {
	sa1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
		},
		ltk: true,
	}

	sa2 := secretAgent{
		person: person{
			"Miss",
			"Moneypenny",
		},
		ltk: true,
	}

	p1 := person{
		first: "Dr.",
		last:  "Yes",
	}

	fmt.Println(sa1)
	sa1.speak()
	sa2.speak()

	fmt.Println(p1)
	p1.speak()
	bar(sa1) //since both secret agent and person have the method speak() attached, they are also of TYPE human and can be passed as arguments to bar()
	bar(sa2)
	bar(p1)

	//Conversion
	fmt.Println()
	var x hotdog = 42
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	var y int = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)

}
