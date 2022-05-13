package main

import "fmt"

func main() {

	//arr := []int{2, 3, 4, 5, 6, 7, 8, 9}

	//x := sum(arr...) //sum accepts ints and NOT []int, so we have to UNBOX the slice

	//fmt.Println("Sum from the main:", x)

	y := sum() //no arguments! the function accepts 0-many arguments

	fmt.Println("sum of y:", y)
	fmt.Println()

	fmt.Println("Another function with 2 params. Regular and Variadic")
	x := sum2("Jose") // no variadic; the underlying array will not be created (nil)
	fmt.Println("total is:", x)
	fmt.Println()

	z := sum2("Raul", 2, 3, 4, 5, 6, 7, 8, 9) //passing the string and variadic; underlying array is created
	fmt.Println("z total is:", z)
	fmt.Println()

	arr := []int{2, 3, 4, 5, 6}

	w := sum2("Isabel", arr...) //the underlying array is NOT created since, the []int is passed unchanged. The underlying array is the one passed.
	fmt.Println("w total is:", w)
}

func sum2(s string, x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x) //[]int -> slice of ints
	fmt.Println(len(x))   //length
	fmt.Println(cap(x))   //cap

	sum := 0

	for _, val := range x {
		sum += val
	}

	return sum
}

func sum(x ...int) int {

	fmt.Println(x)
	fmt.Printf("%T\n", x) //[]int -> slice of ints
	fmt.Println(len(x))
	fmt.Println(cap(x))

	sum := 0

	for _, val := range x {
		sum += val
	}

	return sum
}
