package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	defer foo()
	_, err := os.Open("youcantseeme.txt")
	if err != nil {
		log.Panicln(err)
	}
}

func foo() {
	fmt.Println("WILL I SEE THE SUN!???")
}

/*
	Panic is equivalent to println() followed by a call to panic
*/
