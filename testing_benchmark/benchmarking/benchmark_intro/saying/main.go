// Package saying salutes you
package saying

import "fmt"

// Greet will take a name as a string and return a greeting as string
func Greet(s string) string {
	return fmt.Sprint("Welcome my dear ", s)
}
