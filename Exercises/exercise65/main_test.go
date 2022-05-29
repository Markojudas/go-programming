package main

import (
	"math/rand"
	"testing"
)

func TestFindAverage(t *testing.T) {

	got := FindAverage([]float64{1, 2})

	if got != 1.5 {
		t.Errorf("findavarege(1,2) = %f; expected 1.5", got)
	}

}

//Benchmarks
func BenchmarkFindAverage(b *testing.B) {

	for i := 0; i < b.N; i++ {
		rand.Float64()
	}

}
