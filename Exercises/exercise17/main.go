package main

import "fmt"

//a: show "if statement in action"
//b: show else if, and else

func main() {
	a := 'A'
	//a := 'Z'
	//a := 'x'

	if a < 'Z' {
		fmt.Printf("%c has a smaller ASCII value, %v, than %c (%v)\n", a, a, 'Z', 'Z')
	} else if a == 'Z' {
		fmt.Printf("%c and %c are equal, %v %v\n", a, 'Z', a, 'Z')
	} else {
		fmt.Printf("%c has a greater ASCII value, %v, than %c (%v)\n", a, a, 'Z', 'Z')

	}
}
