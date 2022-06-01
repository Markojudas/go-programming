// Package cavg helps to calculate the average of a given slice of ints.
package cavg

import "sort"

// CenteredAvg accepts a slice of ints and sorts it.
// It then returns the average after removing the smallest and largest values.
func CenteredAvg(xi []int) float64 {
	sort.Ints(xi)

	xi = xi[1 : len(xi)-1]

	total := 0

	for _, val := range xi {
		total += val
	}

	average := float64(total) / float64(len(xi))

	return average
}
