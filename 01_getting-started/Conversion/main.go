package main

import "fmt"

var a int

type hotdog int //new custom type - underlined type is int

var b hotdog

func main() {
	a = 42
	b = 43
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(b)
	fmt.Printf("%T\n", b) //TYPE: main.hotdog
	a = int(b)            //conversion; b is type hotdog but underlined as int; a holds ints
	fmt.Println(a)
	fmt.Printf("%T\n", a)

	// there is no casting but conversion
}
