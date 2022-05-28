package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {

	_, err := sqrt(-10)
	if err != nil {
		log.Fatalln(err)
		fmt.Println("You should log it - fatal!!!") //this doesn't run
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("norgate math: square root of negative number")
	}
	return 42, nil
}
