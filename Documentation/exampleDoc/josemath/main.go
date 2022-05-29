// Package josemath provides an example of documentation and silly math
package josemath

// Sum adds all the elements
func Sum(elem ...int) int {
	sum := 0

	for _, val := range elem {
		sum += val
	}

	return sum
}
