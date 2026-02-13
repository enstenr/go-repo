package main

import (
	"fmt"
	"math"
)

func helper(arr []int, l, h, target int) int {

	if l > h {
		return -1
	}

	m := int(math.Floor(float64((l + h) / 2)))
	if arr[m] == target {
		return m
	}

	if arr[m] < target {
		l = m + 1
		return helper(arr, l, h, target)
	}
	if arr[m] > target {
		h = m - 1

		return helper(arr, l, h, target)
	}
	return m
}
func binarysearch2(arr []int, target int) int {
	return helper(arr, 0, len(arr)-1, target)

}

func main() {

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	target := 3
	fmt.Printf("Found %d at %d ", target, binarysearch2(arr, target))
}
