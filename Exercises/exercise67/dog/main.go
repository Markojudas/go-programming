// Package dog allows us to convert human years to dog years
package dog

// Years converts human years to dog years
func Years(n int) int {
	return n * 7
}

// YearsTwo converts human years to dog years using a for loop
func YearsTwo(n int) int {
	count := 0
	for i := 0; i < n; i++ {
		count += 7
	}

	return count
}
