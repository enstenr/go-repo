package main

import (
	"strings"
	"testing"
)

func BenchmarkReverse(b *testing.B) {
	// 1. Create the same massive test data (multiply by 1000)
	testData := strings.Repeat("gninraeL nIdekniL htiw tol a nraeL", 1000)

	// 2. Setup the struct once outside the loop
	s := Stack{str: testData}

	// 3. Reset the timer so we don't measure the setup above
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		s.reverse()
	}
}
