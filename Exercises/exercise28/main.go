package main

import "fmt"

type person struct {
	first         string
	last          string
	fav_ice_cream []string
}

func main() {

	p1 := person{
		first: "Jose",
		last:  "Hernandez",
		fav_ice_cream: []string{
			"Strawberry",
			"Vanilla",
			"Pineapple",
		},
	}

	p2 := person{
		first: "Damarys",
		last:  "Hernandez",
		fav_ice_cream: []string{
			"Coconut",
			"Chocolate",
			"Cookie Dough",
		},
	}

	fmt.Println("Favorite Ice Cream for", p1.first)
	for _, ice_cream := range p1.fav_ice_cream {
		fmt.Println(ice_cream)
	}

	fmt.Println()

	fmt.Println("Favorite Ice Cream for", p2.first)
	for _, ice_cream := range p2.fav_ice_cream {
		fmt.Println(ice_cream)
	}

}
