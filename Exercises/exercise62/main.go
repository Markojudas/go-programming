package main

import (
	"encoding/json"
	"fmt"
	"log"
)

/*
	Using this code: https://go.dev/play/p/3W69TH4nON
	Create a custom error message using "fmt.Errof()"

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

	bs, err := toJSON(p1)

	if err != nil {
		log.Fatalln("Error:", err)
	}

	fmt.Println(string(bs))

}

func toJSON(a interface{}) ([]byte, error) {
	bs, err := json.Marshal(a)

	if err != nil {
		return []byte{}, fmt.Errorf("input error: %v", err)

	}

	return bs, nil
}
