package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	fp, err := os.Create("Names.txt")

	if err != nil {
		fmt.Println(err)
	}
	defer fp.Close() //always close the file; defer it so it is the last thing it runs after main

	r := strings.NewReader("Jose Hernandez")

	io.Copy(fp, r) //reading from r and write into fp
}
