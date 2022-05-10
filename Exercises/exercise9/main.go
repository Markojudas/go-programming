package main

import "fmt"

//assigns an int to a variable
//prints that int in decimal, binary, and hex
//shifts the bits of that int over 1 position to the left, and assigns that to a variable
//prints the variable in decimal, binary, and hex

func main() {
	a := 42 // assigned int to a variable

	fmt.Printf("%d\t\t%b\t\t%#x\n", a, a, a) //decimal, binary, hex

	b := a << 1 //shifting the bits of a over 1 position to the left ; 74

	fmt.Printf("%d\t\t%b\t\t%#x\n", b, b, b) //decimal, binary, hex

}
