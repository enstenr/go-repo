package main

import "fmt"

func target_sum_efficient(arr []int, target int) (int, int) {
	first := 0
	last := len(arr) - 1
	for first < last {
		sum := arr[first] + arr[last]
		if sum == target {
			return arr[first], arr[last]
		} else if sum < target {
			first++
		} else if sum > target {
			last--
		}
	}
	return 0, 0
}
func target_sum(nums []int, target int) (int, int) {
	i, j := 0, 1
	for i = 0; i < len(nums); i++ {
		for j = i + 1; j < len(nums); {
			if nums[i]+nums[j] == target {
				return nums[i], nums[j]
			}
			j++
		}

	}
	return 0, 0
}
func main() {
	nums := []int{2, 7, 11, 15}
	target := 26
	fmt.Println(target_sum(nums, target))

	fmt.Println(target_sum_efficient(nums, target))
}
