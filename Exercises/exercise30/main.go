package main

import "fmt"

//Creating type: Vehicle
type vehicle struct {
	doors int
	color string
}

//creating type: truck
type truck struct {
	vehicle
	fourWheel bool
}

//creating type: sedan
type sedan struct {
	vehicle
	luxury bool
}

func main() {

	v1 := truck{
		vehicle: vehicle{
			doors: 4,
			color: "red",
		},
		fourWheel: true,
	}

	v2 := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "black",
		},
		luxury: false,
	}

	fmt.Println(v1, v2)

	fmt.Println(v1.fourWheel)
	fmt.Println(v2.luxury)

}
