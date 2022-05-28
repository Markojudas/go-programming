package main

import (
	"encoding/json"
	"fmt"
	"log"
)

/*
	Using this code: https://go.dev/play/p/3W69TH4nON
	Instead of using the blank identifier, make sure the code is checking and handling the error

*/

type person struct {
	First   string
	Last    string
	Sayings []string
}

func main() {

	p1 := person{
		First: "James",
		Last:  "Bond",
		Sayings: []string{
			"Shaken, not stirred",
			"Any last wishes",
			"Never say never",
		},
	}

	bs, err := json.Marshal(p1)

	if err != nil {
		log.Fatalln("Erro:", err)
	}
	fmt.Println(string(bs))

}
