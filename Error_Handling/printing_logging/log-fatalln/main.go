package main

import (
	"fmt"
	"log"
	"os"
)

// the Fatal functions calls os.Exit(1)
// log.Fatalln writes to log and calls os.Exit(1)

//when fatal occurs any deferred functions DON'T RUN

func main() {

	defer foo()

	_, err := os.Open("no file.txt")

	if err != nil {
		log.Fatalln(err)
	}
}

func foo() {
	fmt.Println("YOU WONT'T SEE ME AS I WON'T PRINT")
}
