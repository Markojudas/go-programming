package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func changeMe(p *person, s string) {
	p.first = s
	//(*p).first = s -> equivalent to above
}

func main() {
	p1 := person{
		first: "Jose",
		last:  "Hernandez",
		age:   37,
	}
	fmt.Println(p1.first)

	changeMe(&p1, "Markojudas")

	fmt.Println("Has now become", p1.first, "!")
}
