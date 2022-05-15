package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

type square struct {
	sides float64
}

type shape interface {
	area() float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (s *square) area() float64 {
	return s.sides * s.sides
}

func info(s shape) {
	fmt.Println("Area:", s.area())
}

func main() {

	//creating a circle
	cir := circle{5}

	//passing a non-pointer value to a method with a non-pointer receiver
	//func (c circle) area() float64{ //// }
	info(cir)

	//it also works passing a pointer
	info(&cir)

	//creating a square
	sq := square{10}

	//The area function of square has a POINTER receiver; only works with a POINTER
	info(&sq) //100

	//It does not work passing a NON-POINTER
	//info(sq) --> will not compile.

}
