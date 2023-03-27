package main

import (
	"testing"
)

func Test_Count(t *testing.T) {
}

func Benchmark_Count(b *testing.B) {
	for i := 0; i < b.N; i++ {
	}
}
