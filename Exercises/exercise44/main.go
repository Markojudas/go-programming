package main

/*
	Marshall the []User to JSON
*/

import (
	"encoding/json"
	"fmt"
)

type User struct {
	First string
	Age   int
}

func main() {

	u1 := User{
		First: "James",
		Age:   32,
	}

	u2 := User{
		First: "Moneypenny",
		Age:   27,
	}

	u3 := User{
		First: "M",
		Age:   54,
	}

	users := []User{u1, u2, u3}

	fmt.Println(users)

	//Here it comes
	output, err := json.Marshal(users)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(string(output))
	}

}
