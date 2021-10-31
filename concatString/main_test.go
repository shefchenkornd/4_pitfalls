package main

import "testing"

// 65 ns/op
func Benchmark_cBytes(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		cBytes(params)
	}

	b.StopTimer()
}

// 99 ns/op
func Benchmark_cPlus(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		cPlus(params)
	}

	b.StopTimer()
}

// 78 ns/op
func Benchmark_cStrings(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		cStrings(params)
	}

	b.StopTimer()
}


// 50 ns/op
func Benchmark_cCopy(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		cCopy(params)
	}

	b.StopTimer()
}
