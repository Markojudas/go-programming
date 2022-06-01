package dog

import (
	"fmt"
	"testing"
)

func TestYears(t *testing.T) {

	human_years := 10

	dog_years := Years(human_years) //259

	if dog_years != 70 {
		t.Error("Expected:", 70, "Got:", dog_years)
	}
}

func TestYearsTwo(t *testing.T) {

	human_years := 10

	dog_years := Years(human_years) //259

	if dog_years != 70 {
		t.Error("Expected:", 70, "Got:", dog_years)
	}
}

func ExampleYears() {

	fmt.Println(Years(10))
	// Output:
	// 70
}

func ExampleYearsTwo() {

	fmt.Println(YearsTwo(10))
	// Output:
	// 70
}

func BenchmarkYears(b *testing.B) {

	for i := 0; i < b.N; i++ {
		Years(10)
	}
}

func BenchmarkYearsTwo(b *testing.B) {

	for i := 0; i < b.N; i++ {
		YearsTwo(10)
	}
}
