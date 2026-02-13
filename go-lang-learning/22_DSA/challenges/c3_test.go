package main

import "testing"

// go test -bench=. -benchmem c3.go c3_test.go
// function name must start with Benchmark
func BenchmarkC3Testing1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		target_sum([]int{11, 22, 3, 34, 5, 33, 12, 44, 54, 44, 99, 12, 15, 17, 81}, 100)
	}
}

func BenchmarkC3Testing2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		target_sum_efficient([]int{11, 22, 3, 34, 5, 33, 12, 44, 54, 44, 99, 12, 15, 17, 81}, 100)
	}
}
