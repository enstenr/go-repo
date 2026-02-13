package main

import "fmt"

func binarysearch(nums []int, target int) int {

	l := 0
	h := len(nums) - 1

	for l <= h {

		mid := (l + h) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			h = mid - 1
		}
		if nums[mid] < target {
			l = mid + 1
		}
	}

	return -1

}
func main() {
	arra := []int{1, 3, 5, 7, 9}
	target := 9
	fmt.Printf("Found %d at %d ", target, binarysearch(arra, target))
}
