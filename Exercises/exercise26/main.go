package main

import "fmt"

/*
	Create a slice of a slice of string. Store the following:
		* "James" "Bond" "Shaken, not stirred"
		* "Miss" "Moneypenny"m "Hellooooo James"
	Range over the records, then range over the data i each record
*/
func main() {

	arr := [][]string{
		{"James", "Bond", "Shaken, not stired"},
		{"Miss", "Monneypenny", "Hellooo, James"},
	}

	for _, record := range arr {
		fmt.Println(record)
		for i, value := range record {
			fmt.Printf("\tvalue %d: %v\n", i+1, value)
		}
		fmt.Println()
	}
}
