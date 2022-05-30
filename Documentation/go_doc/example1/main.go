package main

import (
	"fmt"

	"github.com/markojudas/go-programming/Documentation/go_doc/example1/mymath"
)

func main() {

	fmt.Println("2 + 3 =", mymath.Sum(2, 3))
	fmt.Println("3 + 4 =", mymath.Sum(3, 4))
	fmt.Println("5 + 9 =", mymath.Sum(5, 9))

}
