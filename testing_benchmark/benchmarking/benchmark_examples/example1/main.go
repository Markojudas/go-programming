package main

import (
	"fmt"
	"strings"

	"github.com/markojudas/go-programming/testing_benchmark/benchmarking/benchmark_examples/example1/mystr"
)

const s = "Each step I take, may it hurt may it ache, leads me further away from the past. But as long as I breathe, " +
	"each smile in my bleak face, I'm on my way to find back to the peace of mind"

func main() {

	xs := strings.Split(s, " ") //splitting the string using white space as delim or separateor. Returns a slice of strings []string

	//print each string created by splitting the s.
	for _, val := range xs {
		fmt.Println(val)
	}

	fmt.Printf("\n%s\n", mystr.Cat(xs))
	fmt.Printf("\n%s\n", mystr.Join(xs))
}
