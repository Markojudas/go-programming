package main

import "fmt"

// print out the remainder (%; modulo) which is found for each number between 10 and 100 when divided by 4

func main() {

	for i := 10; i <= 100; i++ {
		fmt.Println(i % 4)
	}
}
