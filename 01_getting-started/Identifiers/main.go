package main

import "fmt"

func main() {

	x := 42

	//short declaration (:=) declares a variable AND assigns a value

	fmt.Println(x)

	//now if you want to reassign... since the variable was already declared!
	x = 69

	fmt.Println(x)

	y := 100 + 24 //statement, made of expressions

	fmt.Println(y)

}
