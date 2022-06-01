package main

import (
	"fmt"

	"github.com/markojudas/go-programming/Exercises/exercise68/quote"
	"github.com/markojudas/go-programming/Exercises/exercise68/word"
)

func main() {

	fmt.Println(word.Count(quote.SunAlso))

	for key, val := range word.UseCount(quote.SunAlso) {
		fmt.Println("Word:", key, "\tCount:", val)
	}
}
