package main

import (
	"math/rand"
	"testing"
)

func Test_Count(t *testing.T) {
}

func Benchmark_Rand(b *testing.B) {
	r := rand.New(NewLockSource())
	for i := 0; i < b.N; i++ {
		r.Int63()
	}
}

func Benchmark_RandPackage(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rand.Int63()
	}
}
