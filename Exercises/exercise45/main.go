package main

import (
	"encoding/json"
	"fmt"
)

/*
	unmarshall the JSON into a Go data structure
*/

type Person struct {
	First   string   `json:"First"`
	Last    string   `json:"Last"`
	Age     int      `json:"Age"`
	Sayings []string `json:"Sayings"`
}

func main() {

	str := `[{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},{"First":"M","Last":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}]`

	// The JSON
	byteSliceStr := []byte(str)

	// Where the Data will be stored
	var people []Person

	err := json.Unmarshal(byteSliceStr, &people)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("data:", people)

		fmt.Println()

		for _, person := range people {
			fmt.Println(person.First, person.Last)
			fmt.Println("Age:", person.Age)
			fmt.Println("Sayings:")
			for _, saying := range person.Sayings {
				fmt.Println("\t", saying)
			}
			fmt.Println()
		}
	}

}
