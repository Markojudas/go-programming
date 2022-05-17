package main

import (
	"fmt"
	"sort"
)

type User struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type ByAge []User
type ByLastName []User

//Attaching these functions to the Type ByAge to implement the sort.Interface interface
func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

//Attaching these functons to the TYPE ByLastName to implement th sort.Interface interface
func (ln ByLastName) Len() int           { return len(ln) }
func (ln ByLastName) Swap(i, j int)      { ln[i], ln[j] = ln[j], ln[i] }
func (ln ByLastName) Less(i, j int) bool { return ln[i].Last < ln[j].Last }

func main() {
	u1 := User{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := User{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := User{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []User{u1, u2, u3}

	fmt.Println("Original Slice")
	pleasantPrint(users)

	sort.Sort(ByAge(users))
	fmt.Println("\nSorted By Age")
	pleasantPrint(users)

	sort.Sort(ByLastName(users))
	fmt.Println("\nSorted By Last Name")
	pleasantPrint(users)

	// Sorting the sayings!
	fmt.Println("\nSorting the Sayings for Each User")
	sortSayingsAndPrint(users)

}

func pleasantPrint(users []User) {

	fmt.Println("*********************************************************")
	for _, user := range users {
		fmt.Println("User:", user.First, user.Last)
		fmt.Println("Age:", user.Age)
		fmt.Println("Sayings:")
		for _, saying := range user.Sayings {
			fmt.Println("\t", saying)
		}
		fmt.Println("-----------------------------------------------------------")
		fmt.Println()
	}
	fmt.Println("**********************************************************")

}

func sortSayingsAndPrint(users []User) {

	for _, user := range users {
		sort.Strings((user.Sayings))
	}
	pleasantPrint(users)
}
