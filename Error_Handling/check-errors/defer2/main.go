package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	fp, err := os.Open("Names.txt")

	if err != nil {
		fmt.Println(err)
	}

	defer fp.Close()

	//reads from a file until error or EOF and returns the data read, []byte, err
	//If successful err == nil and NOT EOF since is not considered an error
	bs, err := io.ReadAll(fp)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(bs))
}
