package main

import "fmt"

/*
	Always check for errors. Maybe not with fmt.Println

	fmt.Print returns the number of bytes written (int) and an error

*/

func main() {

	n, err := fmt.Println("Hello")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(n)   // 6, including the \n
	fmt.Println(err) // <nil>

}
