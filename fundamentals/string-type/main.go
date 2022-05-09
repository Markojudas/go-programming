package main

import "fmt"

func main() {
	s := "Hello!"
	bts := `you can use back-ticks!,
	this is a raw string literal so it includes new line and all!`

	fmt.Println(s)
	fmt.Printf("%T\n", s)
	fmt.Println(bts)
	fmt.Printf("%T\n", bts)

	//a sring is a slice of bytes - a sequence of bytes
	bs := []byte(s) // type conversion, since the underlying type of string is []byte
	fmt.Println(bs) // will printout the ascii value of each character
	fmt.Printf("%T\n", bs)

	fmt.Println("The UTF-8 Code")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U ", s[i])
	}

	fmt.Println()

	s = "Hello World!"

	for i, v := range s {
		fmt.Printf("At Index position %d we have hex %#x --> %c\n", i, v, v)
	}

}
