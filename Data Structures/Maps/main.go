package main

import "fmt"

//Maps are key-value stores for super fast look-up
//unordered list

func main() {
	//map[key]values : composite literal
	m := map[string]int{
		"James":           32,
		"Miss Moneypenny": 27,
	}

	fmt.Println(m)

	fmt.Println(m["James"])

	//if the key doesn't match; you get a 0
	fmt.Println(m["Jose"])

	//check if there is an entry in that map matching the passed key, use OK
	val, ok := m["Jose"]
	fmt.Println(val)
	fmt.Println(ok) //false

	//comma OK idiom
	if val, ok := m["James"]; ok {
		// ; ok : ok == true
		fmt.Println("THIS IS THE IF PRINT", val)
	}

	//HOW TO ADD ELEMENT TO THE MAP

	m["Todd"] = 33 //map[new_key] = value

	//RANGE
	for key, val := range m {
		fmt.Println(key, val)
	}

	//DELETE
	//delete(<map name>, "key")

	delete(m, "James")
	fmt.Println(m)

	//you can delete a record that doesn't exist!
	delete(m, "Jose")
	fmt.Println(m)

	//use the comma ok to check if it exist before deleting
	if val, ok := m["Miss Moneypenny"]; ok {
		fmt.Println("Value: ", val)
		delete(m, "Miss Moneypenny")
	}

	if _, ok := m["Jose"]; !ok {
		fmt.Println("Jose is not found")
	}

	fmt.Println(m)

	if _, ok := m["Todd"]; ok {
		fmt.Println("Deleting Todd")
		delete(m, "Todd")
	}

	fmt.Println(m)
}
