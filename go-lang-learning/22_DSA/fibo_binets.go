package main

import (
	"fmt"
	"math"
)

func fiboO1(n int) int {
	phi := (1 + math.Sqrt(5)) / 2

	// We only need the first part of the formula and round it
	// because the second part (psi) becomes negligible as n grows.
	result := math.Round(math.Pow(phi, float64(n)) / math.Sqrt(5))

	return int(result)
}

func main() {
	fmt.Println(fiboO1(10)) // Output: 55
}
