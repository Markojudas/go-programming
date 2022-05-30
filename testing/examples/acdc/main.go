// Package acdc asks you if you are ready to rock
package acdc

// Sum adds an unlimited number of values of type int
func Sum(n ...int) (sum int) {

	sum = 0

	for _, val := range n {
		sum += val
	}

	return sum
}
