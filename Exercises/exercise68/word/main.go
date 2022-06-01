// Package word calculates both the word count and the amount of times each word from the string is used.
package word

import "strings"

// UseCount returns a map storing each word as a key and the number of its occurrances in the string.
func UseCount(s string) map[string]int {
	xs := strings.Fields(s)

	m := make(map[string]int)

	for _, val := range xs {
		m[val]++
	}

	return m
}

// Count returns as an int the word count.
func Count(s string) int {
	xs := strings.Fields(s)

	return len(xs)
}
