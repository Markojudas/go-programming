// Package mjmath provides ACME inc math solutions
package mjmath

// Sum adds an unlimited number of values of type int
func Sum(elem ...int) int {
	sum := 0

	for _, val := range elem {
		sum += val
	}

	return sum
}
