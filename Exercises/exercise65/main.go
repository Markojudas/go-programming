package main

import "fmt"

/*
	Testing and Benchmarking:
	read the docs:https://pkg.go.dev/testing?utm_source=godoc


*/

func main() {

	xs := []float64{1, 2}

	fmt.Println("The average of:", xs)
	fmt.Println(FindAverage(xs))

}

//function to test
func FindAverage(xs []float64) float64 {
	total := float64(0)

	//sum of all values in xs
	for _, val := range xs {
		total += val
	}

	//returns the average by dividing it by the number of elements
	return total / float64(len(xs))
}
