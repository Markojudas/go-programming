package main

import "fmt"

func main() {
	/* switch { a missing switch expression is equivalent to the boolean value "true".
	case false:
		fmt.Println("This will not print")
	case (2 == 4):
		fmt.Println("This will not print either")
	case (3 == 3):
		fmt.Println("This will print!")
		fallthrough //everything after prints // DO NOT USE
	case (4 == 4):
		fmt.Println("True, but will it print? hint: needs fallthrough to print")
		fallthrough
	case (7 == 9):
		fmt.Println("THis SHOULDN'T print but it does because of fallthrough")
		fallthrough
	case (15 == 15):
		fmt.Println("this will print becaus it is true BUT also because of fallthrough")
		fallthrough
	default:
		fmt.Println("This shouldn't print since there are TRUE's but it does due to fallthrough")
	} */
	/*
		n := "Bond"
		switch n {
		case "Moneypenny":
			fmt.Println("Miss money")
		case "Bond":
			fmt.Println("Bond, James")
		case "Q":
			fmt.Println("This is is Q")
		default:
			fmt.Println("This is default")
		} */

	// same example as above but with multiple cases

	n := "Bond"
	switch n {
	case "Moneypenny", "Bond", "Do No":
		fmt.Println("Miss money or bond or dr no")
	case "M":
		fmt.Println("Bond, James")
	case "Q":
		fmt.Println("This is is Q")
	default:
		fmt.Println("This is default")
	}

}
