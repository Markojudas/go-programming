package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	// creating a log file to... log? any errors!
	fp, err := os.Create("Log.txt")

	if err != nil {
		fmt.Println(err)
	}
	defer fp.Close()

	log.SetOutput(fp)

	// forcing an error to be logged
	fp2, err := os.Open("ghostfile.txt")

	if err != nil {
		log.Println("Error Happened:", err) //the output WILL be logged
	}
	defer fp2.Close()

	fmt.Println("Check the log.txt file in the directory")
}
