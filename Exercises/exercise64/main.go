package main

import (
	"fmt"
	"log"
)

/*
	Using this code:https://go.dev/play/p/wlEM1tgfQD

	Use the sqrt.Error struct as a value of type error.


*/

type sqrtError struct {
	lat  string
	long string
	err  error
}

func (se sqrtError) Error() string {
	return fmt.Sprintf("math error: %v %v %v", se.lat, se.long, se.err)
}

func main() {

	_, err := sqrt(-10.23)
	if err != nil {
		log.Println(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		//write code here
		errorCode := sqrtError{
			lat:  "50.2289 N",
			long: "99.4656 W",
			err:  fmt.Errorf("square root of a negative number - value passed: %v", f),
		}
		return 0, errorCode
	}
	return 42, nil
}
