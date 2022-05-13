package main

import "fmt"

//Storing the values we created on exercise 28 on a map and then ranging through the values
//the key of the map is the last name, value is the actual value (akin to object)

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
		last:  "Venegas",
		fav_ice_cream: []string{
			"Coconut",
			"Chocolate",
			"Cookie Dough",
		},
	}

	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	//fmt.Println(m)

	for _, val := range m {
		fmt.Println(val.first)
		fmt.Println(val.last)
		fmt.Println("Fav Ice Cream Flavours:")
		for _, ice_cream := range val.fav_ice_cream {
			fmt.Println(ice_cream)
		}
		fmt.Println()
	}

}
