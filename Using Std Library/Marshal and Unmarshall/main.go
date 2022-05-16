package main

import (
	"encoding/json"
	"fmt"
)

//marshall example
// Used to send data in JSON; returns a []byte (bs, byteslice) and err (an error)
// if err is NOT nil then there was an error
//json.Marshall (x interface{}) ([]bytes, error)// it accepts ANY type and returns []bytes and error
//package: encoding/json
//Marshall returns JSON encoded

//Unmarshall
// It gets a JSON file and puts it back into GO
//func Unmarshall(data []byte, v interface{}) error; takes in byte slice and stores it into whatever v is pointing to (aka v is a pointer to any type)

type person struct {
	first string //this wont work for Marshalling, the identifier needs to be capitalized
	last  string
	age   int
}

type person2 struct {
	First string
	Last  string
	Age   int
}

//for unmarshalling we need a structure
type personUnmarshall struct {
	First string `json:"First"` //need the tags for mapping
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

func main() {

	p1 := person{
		first: "James",
		last:  "Bond",
		age:   32,
	}

	p2 := person{
		first: "Miss",
		last:  "Moneypenny",
		age:   27,
	}
	p1_2 := person2{
		First: "James",
		Last:  "Bond",
		Age:   32,
	}

	p2_2 := person2{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
	}
	people := []person{p1, p2}

	fmt.Println(people)

	// Send this data? Marshall -> json

	bs, err := json.Marshal(people)

	//err checking
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(string(bs)) //this will be empty -> change the fields uppercase so it can be exported

	//same example but using person2 which has the fields uppercase

	people2 := []person2{p1_2, p2_2}
	fmt.Println()
	fmt.Println("With capitalized Fields")
	bs2, err2 := json.Marshal(people2)

	if err2 != nil {
		fmt.Println("Error:", err2)
	}
	fmt.Println(string(bs2)) //this will work!!!
	//output:[{"First":"James","Last":"Bond","Age":32},{"First":"Miss","Last":"Moneypenny","Age":27}]
	fmt.Println()

	//UNMARSHALL
	fmt.Println("Grabbing the JSON encoded data for Unmarshalling")

	s := `[{"First":"James","Last":"Bond","Age":32},{"First":"Miss","Last":"Moneypenny","Age":27}]`

	bs3 := []byte(s) //conversion to []bytes

	//creating the interface where the data will be store

	var people3 []personUnmarshall
	//or
	//people3 := []personUnmarshall{} //empty since it will automatically added

	err3 := json.Unmarshal(bs3, &people3) //need to pass the address!

	//error checking
	if err3 != nil {
		fmt.Println(err3)
	}
	fmt.Println("all of the data", people3)

	//or
	fmt.Println("\nOr we could range over it")
	for idx, val := range people3 {
		fmt.Println("\nPerson NUMBER:", idx)
		fmt.Println(val.First, val.Last, val.Age)
	}
}
