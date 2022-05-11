package main

import "fmt"

func main() {

	for i := 33; i <= 122; i++ {
		fmt.Printf("Decimal number: %v ==> %#U\n", i, i)
	}
}
