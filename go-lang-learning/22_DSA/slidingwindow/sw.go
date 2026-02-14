package main

import "fmt"

// Max Sum Sub-array of Size K
func sliding_window(arr []int, windowSize int) int {
	sum := 0

	for i := 0; i < windowSize; i++ {
		sum += arr[i]
	}
	currentSum := sum
	maxSum := sum

	for i := windowSize; i < len(arr); i++ {
		currentSum = currentSum - arr[i-windowSize] + arr[i]
		if currentSum >= sum {
			maxSum = currentSum
		}
	}
	return maxSum
}

// Brute Force: Re-sums everything every time. O(n*k)
func brute_force_window(arr []int, k int) int {
	maxSum := 0
	for i := 0; i <= len(arr)-k; i++ {
		currentSum := 0
		for j := i; j < i+k; j++ {
			currentSum += arr[j]
		}
		if currentSum > maxSum {
			maxSum = currentSum
		}
	}
	return maxSum
}

func main() {
	arr := []int{100, 100, 1, 1, 1, 100, 101}
	windowSize := 2
	fmt.Println(sliding_window(arr, windowSize))
}
