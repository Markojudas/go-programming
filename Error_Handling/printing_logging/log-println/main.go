package main

import (
	"log"
	"os"
)

func main() {

	_, err := os.Open("whatchuwant.jpg")

	if err != nil {
		log.Println("Error Happened:", err) //this printed the date, time (24-h) stamp of the error!
	}
}
