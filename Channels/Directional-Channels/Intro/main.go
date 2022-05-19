package main

import "fmt"

/*
	Directional Channels
*/

func main() {
	c := make(chan int)    //biderectional
	cr := make(<-chan int) //receive-only
	cs := make(chan<- int) //send-only

	fmt.Printf("%T\n", c)
	fmt.Printf("%T\n", cr)
	fmt.Printf("%T\n", cs)

	//Conversions:
	//General to Specific works but specific to general or specific to specific does not
	fmt.Printf("%T\n", (<-chan int)(c)) // this works since we are converting biderectional (general) to receive-only (specific)
	fmt.Printf("%T\n", (chan<- int)(c)) //this works since we are converting biderectional (general) to send-only (specific)
	//fmt.Printf("%T\n", (chan int)(cr)) this doesn't work since we are trying to convert receive-only (specific) to bidrectional (general)
	//fmt.Printf("%T\n", (chan int)(cs)) this doesn't work since we are trying to convert send-only (specific) to biderectional (general)
	//fmt.Printf("%T\n", (<-chan int)(cs)) This doesn't work since we are trying to convert send-only (specific) to receive-only (specific)
	//fmt.Printf("%T\n", (chan <- int)(cr)) This doesn't work since we are trying to convert receive-only (specific) to send-only (specific)

}
