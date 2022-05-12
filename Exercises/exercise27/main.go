package main

import "fmt"

/*
	Create a map with a key of TYPE string which is a person's "last_first" name, and a value of type []string which stores their favorite
	things. Store three records in your map. Print out all of the values, along with their index position in the slice

*/

func main() {

	characters := map[string][]string{
		"bond_james":      {"Shaken, not stirred", "Martinis", "Women"},
		"moneypenny_miss": {"James Bond", "Literature", "Computer Science"},
		"no_dr":           {"Being evil", "Ice Cream", "Sunsets"},
	}

	for key, values := range characters {
		fmt.Println(key)
		for idx, value := range values {
			fmt.Printf("\tValue at index %d: %v\n", idx, value)
		}
	}
	fmt.Println()

	// Add a record
	characters["hernandez_jose"] = []string{
		"Guitar",
		"Scotch",
		"learning",
	}
	//print the map using range loop
	for character, value := range characters {
		fmt.Println(character, value)
	}
	fmt.Println()

	//Deleting James Bond
	delete(characters, "bond_james")

	for character, value := range characters {
		fmt.Println(character, value)
	}

}
