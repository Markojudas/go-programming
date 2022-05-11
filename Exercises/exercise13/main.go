package main

import "fmt"

//print every rune code point of the uppercase alphabet three times.

func main() {
	for i := 'A'; i <= 'Z'; i++ {
		fmt.Println(i)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t%#U\n", i)
		}
	}
}
