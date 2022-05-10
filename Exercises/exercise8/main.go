package main

import "fmt"

const (
	a = 3
	b = 45.34
	c = "ha!"
)

const (
	d int     = 42
	e float64 = 42.42
	f string  = "haha!"
)

func main() {
	fmt.Println(a, b, c, d, e, f)
}
