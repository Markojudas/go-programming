package main

import "fmt"

/*
using SLICING to recreate the given output

[42 43 44 45 46]
[47 48 49 50 51]
[44 45 46 47 48]
[43 44 45 46 47]
*/

func main() {

	arr := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	i := 0
	for i <= (len(arr) - 5) {
		fmt.Println(arr[i:(i + 5)])
		i++
	}

	//or
	fmt.Println()
	fmt.Println("Just like the exercise")
	fmt.Println(arr[:5])
	fmt.Println(arr[5:])
	fmt.Println(arr[2:7])
	fmt.Println(arr[1:6])

}
