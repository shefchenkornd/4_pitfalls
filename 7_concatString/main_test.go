package main

import "testing"

// чтобы запустить бенчмарк используй команду: `go test -bench=.`

// 92.66 ns/op
func Benchmark_cBytes(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		cBytes(params)
	}

	b.StopTimer()
}

// 129.7 ns/op
func Benchmark_cPlus(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		cPlus(params)
	}

	b.StopTimer()
}

// 100.3 ns/op
func Benchmark_cStrings(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		cStrings(params)
	}

	b.StopTimer()
}


// 67.57 ns/op
func Benchmark_cCopy(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		cCopy(params)
	}

	b.StopTimer()
}
