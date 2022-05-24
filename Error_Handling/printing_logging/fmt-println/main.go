package main

import (
	"fmt"
	"os"
)

func main() {

	_, err := os.Open("No-file.txt") //this will return err != nil

	if err != nil {
		fmt.Println("Error Happened:", err)
	}

}
