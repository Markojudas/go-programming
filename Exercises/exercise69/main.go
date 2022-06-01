package main

import (
	"fmt"

	"github.com/markojudas/go-programming/Exercises/exercise69/cavg"
)

func main() {

	xxi := gen()

	for _, val := range xxi {
		fmt.Println(cavg.CenteredAvg(val))
		fmt.Println()
	}

}

func gen() [][]int {
	a := []int{1, 4, 6, 8, 100}
	b := []int{0, 8, 10, 1000}
	c := []int{9000, 4, 10, 8, 6, 12}
	d := []int{123, 744, 140, 200}

	e := [][]int{a, b, c, d}

	return e
}
