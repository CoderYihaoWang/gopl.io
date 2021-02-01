package popcount

import (
	"math/rand"
	"testing"
	"time"

	gopl "gopl.io/popcount"
	ex3 "popcount/ex3"
	ex4 "popcount/ex4"
	ex5 "popcount/ex5"
)

var data = [1000]uint64{}

func init() {
	rand.Seed(time.Now().UnixNano())
	for i := range data {
		data[i] = rand.Uint64()
	}
}

func BenchmarkEx3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, n := range data {
			ex3.PopCount(n)
		}
	}
}

func BenchmarkEx4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, n := range data {
			ex4.PopCount(n)
		}
	}
}

func BenchmarkEx5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, n := range data {
			ex5.PopCount(n)
		}
	}
}

func BenchmarkGopl(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, n := range data {
			gopl.PopCount(n)
		}
	}
}
