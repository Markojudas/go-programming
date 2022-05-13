package main

import "fmt"

func main() {

	//this is anonymous : the struct doesn't have a name. To keep the code clean. if you only need the struct for a small area.
	p1 := struct {
		first string
		last  string
		age   int
	}{
		first: "James",
		last:  "Bond",
		age:   32,
	}
	fmt.Println(p1)
}
