package mjmath

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {

	type test struct {
		data   []int
		answer int
	}

	tests := []test{
		{[]int{21, 21}, 42},
		{[]int{3, 4, 5}, 12},
		{[]int{1, 1}, 2},
		{[]int{-1, 0, 1}, 0},
	}

	for _, v := range tests {
		x := Sum(v.data...)
		if x != v.answer {
			t.Error("expected", v.answer, "got", x)
		}
	}
}

func ExampleSum() {
	fmt.Println(Sum(2, 3))
}
