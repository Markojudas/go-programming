package main

import "fmt"

var y = 42

func main() {

	//Print = just prints to console
	fmt.Println(y)
	fmt.Printf("%b\n", y)  // Binary
	fmt.Printf("%x\n", y)  // Hex
	fmt.Printf("%#x\n", y) // Hex with prefix 0x

	y = 911

	fmt.Printf("%#x\t%b\t%x\n", y, y, y) // hex w prefix, binary, hex w/o prefix

	//Sprint = returns a string and can be assigned to a variable

	s := fmt.Sprintf("%#x\t%b\t%x\n", y, y, y) //same is assigned to s

	//Fprint = prints to a file!

	fmt.Println(s)        //printing s
	fmt.Printf("%v\n", y) //defaul format
}
