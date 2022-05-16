package main

import (
	"fmt"
	"sort"
)

// sorting a slice int and string
// sort is a package
// Provides primitives for sorting slices and user-defined collections

func main() {
	intArr := []int{4, 7, 3, 42, 99, 18, 16, 56, 12} //unsorted

	strArr := []string{"James", "Q", "M", "Moneypenny", "Dr. No"} //unsorted

	fmt.Println(intArr)
	fmt.Println(strArr)

	fmt.Println()
	fmt.Println("SORT!")

	//sort.Ints()
	//func Ints(a []int) // does not return anything.
	sort.Ints(intArr)
	fmt.Println(intArr)

	//sort.Strings() -> increasing order
	//func.Strings(a []string) //does not return anything.
	sort.Strings(strArr)
	fmt.Println(strArr)
}
