package main

import (
	"fmt"

	"github.com/markojudas/go-programming/Exercises/exercise66/dog"
)

type canine struct {
	name string
	age  int
}

func main() {

	rufus := canine{
		name: "Rufus",
		age:  dog.Years(2),
	}
	fmt.Println(rufus)
}
