package main

import "fmt"

/*
	Using MAKE and APPEND, create a slice to store the names of all the States of USA
	...can't have the underlying array be created more than once
	what is the length, capacity
*/

func main() {

	states := make([]string, 0, 50)

	states = append(states, "Alabama", "Alaska", "Arizona", "Arkansas", "California", "Colorado", "Connecticut",
		"Delaware", "Florida", "Georgia", "Hawaii", "Idaho", "Illinois", "Indiana", "Iowa", "Kansas",
		"Kentucky", "Louisiana", "Maine", "Maryland", "Massachusetts", "Michigan", "Minnesota",
		"Mississippi", "Missouri", "Montana", "Nebraska", "Nevada", "New Hampshire", "New Jersey",
		"New Mexico", "New York", "North Carolina", "North Dakota", "Ohio", "Oklahoma", "Oregon",
		"Pennsylvania", "Rhode Island", "South Carolina", "South Dakota", "Tennessee", "Texas",
		"Utah", "Vermont", "Virginia", "Washington", "West Virginia", "Wisconsin", "Wyoming")

	fmt.Println("The length:", len(states))
	fmt.Println("The capacity:", cap(states))

	for i := 0; i < len(states); i++ {
		fmt.Printf("%d:\t%v\n", i+1, states[i])
	}
}
