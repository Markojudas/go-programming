package main

import "fmt"

/* const a = 42
const b = 42.78
const c = "James Bond" */

//we can also do this so we don't have to write const over and over

const (
	a int     = 42
	b float64 = 42.78
	c string  = "James Bond"
)

// constant typed

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("%T\n", a) //int
	fmt.Printf("%T\n", b) //float64
	fmt.Printf("%T\n", c) //string
}
