package main

import "fmt"

/*
	Get this code running:

	A) https://play.golang.org/p/oB-p3KMiH6
	B) https://play.golang.org/p/_DBRueImEq

*/

func main() {

	cg := make(chan int, 1)

	/* A)
	cs := (chan<- int)(cg) //send only channel!

	go func() {
		cs <- 42
	}()

	fmt.Println(<-cg)

	fmt.Println("-----------------------")
	fmt.Printf("cs\t%T\n", cs)
	*/

	//B)
	cr := (<-chan int)(cg)

	go func() {
		cg <- 42
	}()

	fmt.Println(<-cr)

	fmt.Println("-----------------------")
	fmt.Printf("cr\t%T\n", cr)

}
