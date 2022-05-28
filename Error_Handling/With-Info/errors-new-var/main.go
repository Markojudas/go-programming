package main

import (
	"errors"
	"fmt"
	"log"
)

var errMadeup = errors.New("norgate math: square root of negative number")

func main() {
	//printing the type
	fmt.Printf("%T\n", errMadeup)

	_, err := sqrt(-10)

	if err != nil {
		log.Fatalln(err) //logs and kills the process. If panic you could still recover
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errMadeup
	}
	return 42, nil
}
