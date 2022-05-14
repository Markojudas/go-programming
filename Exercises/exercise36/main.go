package main

import (
	"fmt"
	"math"
)

/*
	Create a TYPE SQUARE
	Create a TYPE CIRCLE
	Attach a method to each that calculates AREA and returns it
		* circle area = pi * radius * 2
		* square area = L * W
	Create a TYPE SHAPE that defines an interface as anything that has the AREA method
	Create a func INFO which takes shape and then prints the area
	Create a value of type SQUARE
	Create a value of type CIRLCE
	use func info to print the area of a square
	use func info to print the area of a circle
*/

type circle struct {
	radius float64
}

type square struct {
	length float64
	width  float64
}

func (c circle) area() float64 {

	return math.Pi * c.radius * c.radius
}

func (s square) area() float64 {

	return s.length * s.width
}

type shape interface {
	area() float64
}

func info(s1 shape) {

	fmt.Println(s1.area())
}

func main() {
	circle1 := circle{
		radius: 12.345,
	}
	square1 := square{
		length: 15,
		width:  15,
	}

	info(circle1)
	info(square1)

}
