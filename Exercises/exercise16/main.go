package main

import "fmt"

// using for{} print out the years you have been alive

func main() {

	bd := 1985

	for {
		if bd > 2022 {
			break
		}
		fmt.Println(bd)
		bd++
	}
}
