package main

import "fmt"

// playing with anonymous structs!

func main() {

	p1 := struct {
		first   string
		friends map[string]int
		hobbies []string
	}{
		first: "Jose",
		friends: map[string]int{
			"Josh":  111,
			"Mikee": 222,
			"Ivan":  333,
		},
		hobbies: []string{
			"Guitar",
			"Music",
			"Eat!",
		},
	}

	fmt.Println(p1.first)

	for friend, number := range p1.friends {
		fmt.Println(friend, number)
	}

	for _, hobby := range p1.hobbies {
		fmt.Println(hobby)
	}
}
