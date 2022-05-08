package main

import "fmt"

var x = 42
var y = "James Bond"
var z = true

func main() {

	s := fmt.Sprintf(`X = %v
Y = %v
Z = %v`, x, y, z)

	fmt.Println("Printing the value of s")
	fmt.Println(s)
}
