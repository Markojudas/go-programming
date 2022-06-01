package word

import (
	"fmt"
	"testing"

	"github.com/markojudas/go-programming/Exercises/exercise68/quote"
)

func BenchmarkUseCount(b *testing.B) {

	for i := 0; i < b.N; i++ {
		UseCount(quote.SunAlso)
	}
}

func BenchmarkCount(b *testing.B) {

	for i := 0; i < b.N; i++ {
		Count(quote.SunAlso)
	}
}

func TestCount(t *testing.T) {

	str := "The doom of the Noldor grows near"

	wc := Count(str)

	if wc != 7 {
		t.Error("Expected:", 7, "Got:", wc)
	}

}

func TestUseCount(t *testing.T) {
	m := UseCount("the doom of the Noldor grows near")

	for key, val := range m {
		switch key {
		case "the":
			if val != 2 {
				t.Error("Expected:", 2, "got:", val)
			}
		case "doom":
			if val != 1 {
				t.Error("Expected:", 1, "got:", val)
			}
		case "of":
			if val != 1 {
				t.Error("Expected:", 1, "got:", val)
			}
		case "Noldor":
			if val != 1 {
				t.Error("Expected:", 1, "got:", val)
			}
		case "grows":
			if val != 1 {
				t.Error("Expected:", 1, "got:", val)
			}
		case "near":
			if val != 1 {
				t.Error("Expected:", 1, "got:", val)
			}
		}
	}
}

func ExampleCount() {

	str := "The doom of the Noldor grows near"

	fmt.Println(Count(str))
	// Output:
	// 7
}
