package cavg

import (
	"fmt"
	"testing"
)

func BenchmarkCenteredAvg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CenteredAvg([]int{0, 1, 5, 8, 3, 1000})
	}
}

func ExampleCenteredAvg() {

	x1 := []int{0, 1, 5, 8, 3, 1000}

	fmt.Println(CenteredAvg(x1))
	// Output:
	// 4.25
}

func TestCenteredAvg(t *testing.T) {

	type test struct {
		data []int
		ans  float64
	}

	gIntS := []test{
		{[]int{1, 4, 6, 8, 100}, 6},
		{[]int{0, 8, 10, 1000}, 9},
		{[]int{9000, 4, 10, 8, 6, 12}, 9},
		{[]int{123, 744, 140, 200}, 170},
	}

	for _, val := range gIntS {
		res := CenteredAvg(val.data)

		if res != val.ans {
			t.Error("Expected:", val.ans, "Got:", res)
		}
	}
}
