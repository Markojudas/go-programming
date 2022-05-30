package main

import (
	"fmt"

	"github.com/markojudas/go-programming/testing_benchmark/benchmarking/benchmark_example/saying"
)

func main() {

	name := "Rufus"

	fmt.Println(saying.Greet(name))
}
