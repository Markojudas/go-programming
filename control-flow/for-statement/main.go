package main

import "fmt"

func main() {

	fmt.Println("***For-Statement with 1 condition; akin to C while loops***")

	x := 1
	for x < 10 {
		// while x is less than 10
		fmt.Println(x)
		x++
	}
	fmt.Println("done.")
	fmt.Printf("\n***Infinite Loop (while true; or C's for(;;)); pretty much the same as the above***\n")

	x = 0 //resetting x

	for {
		if x > 9 {
			break
		}
		fmt.Println(x)
		x++
	}
	fmt.Println("done.")
}
