package main

import (
	"fmt"

	"github.com/markojudas/go-programming/testing_benchmark/examples/acdc"
)

func main() {
	fmt.Println(acdc.Sum(2, 3, 4, 5, 6, 7, 8, 9))
	fmt.Println(acdc.Sum(2, 3))

}
