package main

import (
	"fmt"
	"runtime"
)

var x int8 = -127 // -128 - 127; anything outside of it is an overflow

//byte is an alias to uint8
//rune is a alias int32

//using package runtime

func main() {
	//x = 2.32332  x is declared as an int and it only holds ints
	fmt.Println(x)
	fmt.Printf("%T\n", x) // int

	fmt.Println("\n****RUNTIME.GOOS and RUNTIME.GOARCH****")
	fmt.Println(runtime.GOOS)   //os: Linux
	fmt.Println(runtime.GOARCH) //arch: amd64

}
