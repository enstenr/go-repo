package main

import "testing"

// function name must start with Benchmark
func BenchmarkAlgoQuickSortTesting(b *testing.B) {
	for i := 0; i < b.N; i++ {
		quicksort([]int{11, 22, 3, 34, 5, 33, 12, 44, 54, 44, 99, 12, 15, 17, 81})
	}
}
