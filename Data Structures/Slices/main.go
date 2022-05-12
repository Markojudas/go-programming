package main

import "fmt"

func main() {
	// Composite Literal ---> x := type{values_of_same_type}
	// a SLICE allows you to group VALUES of the same TYPE
	// A composite literal create new values everytime they are evaluated; used for arrays, slice, maps

	x := []int{4, 5, 7, 8, 42} // a SLICE
	fmt.Println(x)
	/* fmt.Println()
	fmt.Println(len(x)) // length
	fmt.Println()
	fmt.Println(x[0]) //accessing by index position
	fmt.Println(x[1])
	fmt.Println(x[2])
	fmt.Println(x[3])
	fmt.Println(x[4])
	fmt.Println()
	*/
	// Iterate through the SLICE using FOR RANGE

	/* for idx, val := range x {
		fmt.Printf("Index %v: %v\n", idx, val)
	}
	*/
	//Slicin' a Slice!

	/* 	fmt.Println(x[:])   // colon used to slice the slice ; this one gives from [0-len)
	   	fmt.Println(x[1:3]) //index 1 and 2 print

	   	for i, v := range x {
	   		fmt.Println(i, v)
	   	}

	   	//same as above
	   	for i := 0; i < len(x); i++ {
	   		fmt.Println(i, x[i])
	   	}
	*/
	//APPEND A SLICE; appending is a built-in function

	x = append(x, 77, 88, 99, 1014)
	fmt.Println(x)

	y := []int{234, 456, 678, 987}
	x = append(x, y...) //unfurl ; it gets all the value sof Y and appends it
	fmt.Println(x)

	//DELETE FROM A SLICE; use append!

	x = append(x[:2], x[4:]...) //first arg1 is getting a slice; arg2 is getting the ints
	fmt.Println(x)              // deleted index 2 and 3 (no 7 and 8)

	//MAKE; built-in function : if you already know the size

	z := make([]int, 10, 100) //type, size, capacity
	fmt.Println(z)
	fmt.Println(len(z))
	fmt.Println(cap(z)) //capacity
	x[0] = 42
	x[9] = 999
	z = append(z, 3423)
	fmt.Println(z)
	fmt.Println(len(z))
	fmt.Println(cap(z))

	s := make([]int, 10, 12) //TYPE, size, capacity
	s[0] = 22
	s[9] = 33
	s = append(s, 44) //size 11, capacity 12
	s = append(s, 55) //size 12, capacity 12
	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println(cap(s))
	s = append(s, 66) //size 13, capacity 24 -- the underlying Array was deleted and a new one, of double size, size gets created
	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println(cap(s))
	fmt.Println()

	//RANGE
	xi := []int{4, 5, 7, 8, 9, 42}

	for idx, val := range xi {
		fmt.Println(idx, val)
	}

}
