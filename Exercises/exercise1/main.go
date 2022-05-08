package main

import "fmt"

func main() {

	x := 42
	y := "James Bond"
	c := true

	fmt.Println("A) A single print statement")
	fmt.Printf("x = %v\ny = %v\nc = %v\n", x, y, c)

	fmt.Println("\nB) Multiple print statements")
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(c)

}
