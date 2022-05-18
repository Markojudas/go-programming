package main

import (
	"fmt"
	"sort"
)

/*
	Here we are sorting by:
		1st: Age
		2nd: name

	to sort by age we are following the example on godocs.org under the sort package doc
	first create another type that implements the sort.Interface interface; we use sort.Sort(data Interface)

	the sort.Interface:
	type Interface interface {
		Len() int
		Swap(i, j int)
		Less(i, j int)
	}

	for this example we are creating ByAge type:
	type ByAge []person // so we can convert our slice of people to ByAge, which implements sort.Interface and we can call sort.Sort()

	We have to attach the Len, Swap, & Less function to the TYPE ByAge to make sure it implements sort.Interface

	to sort by name is very similar but we are comparing the name field and not the age

*/

type Person struct {
	First string
	Age   int
}

type ByAge []Person
type ByName []Person

//attaching this methods to ByAge type to implement sort.Implement
func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

//attaching this methods to ByName type to implement sort.Implement
func (a ByName) Len() int           { return len(a) }
func (a ByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByName) Less(i, j int) bool { return a[i].First < a[j].First }

func main() {

	p1 := Person{"James", 32}
	p2 := Person{"Moneypenny", 27}
	p3 := Person{"Q", 64}
	p4 := Person{"M", 56}

	people := []Person{p1, p2, p3, p4}

	//unsorted
	fmt.Println(people)

	//let's sort by age
	fmt.Println()
	sort.Sort(ByAge(people)) //passing people, converted to ByAge, to the sort.Sort() //sorts in ascending order
	fmt.Println(people)

	//let's sort by name
	fmt.Println()
	sort.Sort(ByName(people))
	fmt.Println(people)
}
