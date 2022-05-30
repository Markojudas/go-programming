// Package mystr is an example of benchmarking
package mystr

import "strings"

// Cat concatenates the strings from []string and returns it. Not efficient
func Cat(xs []string) string {

	s := "" //empty string

	for idx, val := range xs {
		s += val //adds the string

		//adds a space
		if idx != len(xs)-1 {
			s += " "
		}
	}

	return s
}

// Join does the same but more efficiently
func Join(xs []string) string {
	return strings.Join(xs, " ")
}
