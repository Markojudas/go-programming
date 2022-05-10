package main

import "fmt"

// bit-shifting (<< left; >> right)

const (
	_      = iota             //here the iota starts at 0; we can throw away this variable for this example as it is not used
	kbiota = 1 << (iota * 10) //here iota == 1; one KB is 1 shifted 10 times to the left
	mbiota = 1 << (iota * 10) //here iota == 2; one MB is 1 shifted 20 times to the left
	gbiota = 1 << (iota * 10) //here iota == 3; one GB is 1 shifts 30 times to the left
)

func main() {
	fmt.Println("***EXAMPLE OF SHIFTING****")
	x := 4
	fmt.Printf("%d\t\t%b\n", x, x) //printing decimal and binary of x : 4  0000 0100

	y := x << 1                    // 8
	fmt.Printf("%d\t\t%b\n", y, y) //printing decimal and binary of y : 8  0000 1000

	fmt.Println()
	kb := 1024
	mb := 1024 * kb
	gb := 1024 * mb

	//here the bits are shift ten places each time (1024)
	fmt.Printf("%d\t\t\t%b\n", kb, kb) // 1 shifted 10 times
	fmt.Printf("%d\t\t\t%b\n", mb, mb) // 1 shifted 20 times
	fmt.Printf("%d\t\t%b\n", gb, gb)   // 1 shifted 20 times

	fmt.Println()
	fmt.Println("Now let's do the same example but with IOTA! it should be the same as above")
	fmt.Printf("%d\t\t\t%b\n", kbiota, kbiota)
	fmt.Printf("%d\t\t\t%b\n", mbiota, mbiota)
	fmt.Printf("%d\t\t%b\n", gbiota, gbiota)

}
