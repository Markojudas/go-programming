package main

import "fmt"

const (
	base, fourYears = 2018, base + iota
	threeYears      = base + iota
	twoYears        = base + iota
	oneYear         = base + iota
	thisYear        = base + iota
)

func main() {
	fmt.Println(fourYears, threeYears, twoYears, oneYear, thisYear)
}
