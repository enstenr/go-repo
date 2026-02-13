package main

import "testing"

// function name must start with Benchmark
func BenchmarkBinary1Test(b *testing.B) {
	for i := 0; i < b.N; i++ {
		binarysearch([]int{11, 22, 3, 34, 5, 33, 12, 44, 54, 44, 99, 12, 15, 17, 81}, 88)
	}
}
func BenchmarkBinary2Test(b *testing.B) {
	for i := 0; i < b.N; i++ {
		binarysearch([]int{11, 22, 3, 34, 5, 33, 12, 44, 54, 44, 99, 12, 15, 17, 81}, 81)
	}
}

func BenchmarkBinaryHuge(b *testing.B) {
	// 1. Setup: Create a slice of 1 million elements
	data := make([]int, 1000000)
	for i := 0; i < 1000000; i++ {
		data[i] = i
	}

	// 2. Important: Reset the clock after the setup is done
	b.ResetTimer()

	// 3. The Benchmark loop
	for i := 0; i < b.N; i++ {
		// Search for a number near the end to force the most "steps"
		_ = binarysearch(data, 999999)
	}
}
