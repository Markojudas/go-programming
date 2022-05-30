package main

import "testing"

func TestMySum(t *testing.T) {

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

	for _, val := range tests {
		x := mySum(val.data...)
		if x != val.answer {
			t.Error("expected", val.answer, "got", x)
		}
	}
}
