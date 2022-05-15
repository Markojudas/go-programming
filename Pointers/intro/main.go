package main

import (
	"fmt"
)

// Pointers hold memory addresses

func main() {
	var a int = 5

	fmt.Println(a)         //printing value
	fmt.Println(&a)        //printing the address :: remember that GOLANG is pass by value
	fmt.Printf("%T\n", a)  // TYPE : int
	fmt.Printf("%T\n", &a) // TYPE : *int -> pointer to an int

	/* 	fmt.Println(unsafe.Sizeof(a))
	fmt.Println(unsafe.Sizeof(&a)) */

	//less see what's stored at the address &a:
	b := &a
	fmt.Println(b)  // it will hold the address
	fmt.Println(*b) //dereference --> what is stored in the address stored in B?

	//since be is pointing to the address
	*b = 42
	//now a has been modified; a is the value stored at that address which b modified
	fmt.Println(a)

}
