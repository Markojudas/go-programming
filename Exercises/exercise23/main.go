package main

import "fmt"

/*
	same slice as exercise 22;
	append 52
	print slice
	one statement apped 53, 54, 55
	print slice
	append the slice to the slice
		y := []int{56,57,58,59,60}

	print slice
*/

func main() {

	arrX := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println("arrX")
	fmt.Println(arrX)
	fmt.Println()

	arrX = append(arrX, 52)
	fmt.Println("Appending 52 to arrX")
	fmt.Println(arrX)
	fmt.Println()

	arrX = append(arrX, 53, 54, 55)
	fmt.Println("Appending 53, 54, 55 to the arrX")
	fmt.Println(arrX)
	fmt.Println()

	arrY := []int{56, 57, 58, 59, 60}
	fmt.Println("arrY")
	fmt.Println(arrY)
	fmt.Println()

	fmt.Println("Appending arrY to arrX")
	arrX = append(arrX, arrY...)
	fmt.Println(arrX)

}
