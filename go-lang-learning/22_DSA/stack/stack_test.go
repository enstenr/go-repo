package stack

import (
	"strings"
	"testing"
)

func BenchmarkStackTest(b *testing.B) {
	// 1. Create the string
	testData := strings.Repeat("gninraeL nIdekniL htiw tol a nraeL", 1000)

	// 2. Reset the timer to ignore the string creation time
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		// 3. Call the function that uses Push/Pop logic
		ReverseWithStack(testData)
	}
}
