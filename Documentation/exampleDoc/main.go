package main

import (
	"fmt"

	"github.com/markojudas/go-programming/Documentation/exampleDoc/josemath"
)

func main() {

	fmt.Println("2 + 3 =", josemath.Sum(2, 3))
	fmt.Println("4 + 7 =", josemath.Sum(4, 7))
	fmt.Println("5 + 9 =", josemath.Sum(5, 9))
}
