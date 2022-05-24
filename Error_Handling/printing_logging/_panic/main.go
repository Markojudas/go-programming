package main

import (
	"os"
)

func main() {

	_, err := os.Open("IDisappear.txt")

	if err != nil {
		panic(err)
	}

}
