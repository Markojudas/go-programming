package main

import "fmt"

// using for condition {} print the years you have been alive

func main() {
	birthDay := 1985

	for birthDay <= 2022 {
		fmt.Println(birthDay)
		birthDay++
	}
}
