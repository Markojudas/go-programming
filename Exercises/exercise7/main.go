package main

import "fmt"

// use the boolean operators (==, <=, >=, !=, <, >) to create expressions and assign their values to variables

func main() {
	g := 3 == (2 + 1) // true
	h := 3 <= 4       // true
	i := 3 >= (2 + 1) // true
	j := 3 != (2 + 1) // false
	k := 3 < (2 + 1)  // false
	l := 3 > 2        //true

	fmt.Println(g)
	fmt.Println(h)
	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(k)
	fmt.Println(l)

}
