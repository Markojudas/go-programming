package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

//bcryt is a password hashing function
//golang.org/x/crypto/bcrypt

/*
	To generate a hash/encrypted password:
	func GenerateFromPassword(password []byte, cost int) ([]byte, error)

	takes in the password as a byte slice and the cost (an int) returnning the slice of byte (encrypted password) and an error

	const:
		MinCost = 4
		MaxCost = 31
		DefaultCost = 10; if we pass a cost lower than MinCost

	to check if the passowrd generates the hash (correct password) use:

	func CompareHashAndPassword(hashedPassword, password []byte) error
	it takes the hashed password, and the password, and returns an error; if it doesn't match then it will give an error
*/

func main() {
	s := `password1234`

	hashedPW, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(s)
	fmt.Println(hashedPW)

	loginPW := `password1234`

	err = bcrypt.CompareHashAndPassword(hashedPW, []byte(loginPW))

	if err != nil {
		fmt.Println("Wrong credentials")
	} else {
		fmt.Println("You are in!")
	}

}
