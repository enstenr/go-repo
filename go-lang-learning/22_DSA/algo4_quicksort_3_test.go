package main

import "testing"

// function name must start with Benchmark
func BenchmarkAlgoQuickSort3Testing(b *testing.B) {
	for i := 0; i < b.N; i++ {
		quicksort_4([]int{11, 22, 3, 34, 5, 33, 12, 44, 54, 44, 99, 12, 15, 17, 81})

	}
}
func BenchmarkAlgoQuickSort3TestingSorted(b *testing.B) {
	for i := 0; i < b.N; i++ {
		quicksort_4([]int{3, 5, 11, 12, 12, 15, 17, 22, 33, 34, 44, 44, 54, 81, 99})

	}
}
