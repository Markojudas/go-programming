package mystr

import (
	"fmt"
	"strings"
	"testing"
)

func splitStr(str string) []string {

	return strings.Split(str, " ")
}

func TestCat(t *testing.T) {

	xs := splitStr("Shaken not stirred")

	s := Cat(xs)

	if s != "Shaken not stirred" {
		t.Error("got", s, "want", "Shaken not stirred")
	}
}

func TestJoin(t *testing.T) {

	xs := splitStr("Shaken not stirred")

	s := Join(xs)

	if s != "Shaken not stirred" {
		t.Error("got", s, "want", "Shaken not stirred")
	}
}

func ExampleCat() {

	xs := splitStr("Shaken not stirred")
	fmt.Println(Cat(xs))
	// Output:
	// Shaken not stirred
}

func ExampleJoin() {

	xs := splitStr("Shaken not stirred")
	fmt.Println(Join(xs))
	// Output:
	// Shaken not stirred
}

const s = "Each step I take, may it hurt may it ache, leads me further away from the past. But as long as I breathe, " +
	"each smile in my bleak face, I'm on my way to find back to the peace of mind"

func BenchmarkCat(b *testing.B) {

	xs := strings.Split(s, " ")

	for i := 0; i < b.N; i++ {
		Cat(xs)
	}
}

func BenchmarkJoin(b *testing.B) {

	xs := strings.Split(s, " ")

	for i := 0; i < b.N; i++ {
		Join(xs)
	}

}
