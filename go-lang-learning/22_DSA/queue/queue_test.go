package main

import (
	"strings"
	"testing"
)

//go test -bench=. -benchmem queue.go queue_test.go

func BenchmarkQueueTest(b *testing.B) {
	// 1. Setup data: a slice of strings to pump through the queue
	// Repeating the string to simulate a heavy workload
	testData := strings.Split(strings.Repeat("Spark-Job-ID-", 1000), "-")

	// 2. Reset the timer to ignore the setup
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		q := Queue[string]{items: make([]string, 0, len(testData))}

		// 3. Performance Test: Enqueue all items
		for _, s := range testData {
			q.Enqueue(s)
		}

		// 4. Performance Test: Dequeue all items
		for len(q.items) > 0 {
			_, _ = q.Dequeue()
		}
	}
}
