package main

import "testing"

// function name must start with Benchmark
func BenchmarkStressTestSW(b *testing.B) {
	// Create a large dataset of 10,000 numbers
	data := make([]int, 10000)
	for i := range data {
		data[i] = i % 100
	}
	windowSize := 100

	b.Run("SlidingWindow", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			sliding_window(data, windowSize)
		}
	})

	b.Run("BruteForce", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			brute_force_window(data, windowSize)
		}
	})
}
