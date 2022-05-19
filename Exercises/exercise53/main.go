package main

/*
	*	Run with:
			* go run -> regular running
			* go build -> builds and creates an executable on your folder (exercise53)
			* go install -> creates an executable on $GOPATH/bin/
*/

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("OS:", runtime.GOOS)
	fmt.Println("ARCH:", runtime.GOARCH)
}
