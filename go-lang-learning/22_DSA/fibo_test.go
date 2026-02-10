package main

import "testing"

func BenchmarkFibo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibo(10)
	}
}
func BenchmarkFibo1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibo(1000)
	}
}
