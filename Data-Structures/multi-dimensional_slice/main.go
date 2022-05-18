package main

import "fmt"

func main() {
	jb := []string{"james", "bond", "Chocolate", "Martini"}
	fmt.Println(jb)

	mp := []string{"Miss", "Moneypenny", "Strawbery", "Hazelnut"}
	fmt.Println(mp)

	xp := [][]string{jb, mp} //slice of slices
	fmt.Println(xp)

	fmt.Println(xp[0][2]) //should print "Chocolate"
	fmt.Println(xp[1][3]) //should print "Hazelnut"

}
