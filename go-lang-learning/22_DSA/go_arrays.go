package main

import (
	"fmt"
	"unsafe"
)

func create_array() [3]int16 {
	fmt.Println(unsafe.Sizeof(int(0)))
	arr := [3]int16{1, 2, 3} // array of fixed size 3
	fmt.Printf("%p", &arr)   // prints address to the array
	fmt.Println()
	fmt.Printf("%p", &arr[0]) // prints address to the array
	fmt.Println()
	fmt.Printf("%p", &arr[1]) // prints address to the array
	fmt.Println()
	fmt.Printf("%p", &arr[2]) // prints address to the array
	fmt.Println()
	return arr
}
func create_array_using_make() {
	s := make([]int, 0, 2)
	fmt.Printf("Initial addr = %p, len = %d , cap = %d \n", s, len(s), cap(s))

	for i := 0; i < 5; i++ {
		s = append(s, i)
		fmt.Printf("Initial addr = %p, len = %d , cap = %d \n", s, len(s), cap(s))
	}
}
func reverse(nums []int16) []int16 {
	left := 0
	right := len(nums) - 1
	for left < right {

		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
	return nums
}
func twoSumSorted(nums []int, target int) []int {
	left := 0
	right := len(nums) - 1
	for left < right {
		sum := nums[left] + nums[right]
		if sum == target {
			return []int{nums[left], nums[right]}
		}
		if sum > target {
			right--
		}
		if sum < target {
			left++
		}
	}
	return []int{-1, -1}
}

func removeDuplicates(nums []int) []int {
	left := 0

	for right := 1; right < len(nums); right++ {
		if nums[right] != nums[left] {
			left++
			nums[left] = nums[right]
		}
	}
	return nums
}

func slidingWindow(nums []int, k int) int {
	windowsum := 0
	for i := 0; i < k; i++ {
		windowsum += nums[i]
	}
	maxSum := windowsum

	for i := k; i < len(nums); i++ {
		windowsum += nums[i] - nums[i-k]
		if windowsum > maxSum {
			maxSum = windowsum
		}
	}
	return maxSum
}
func main() {
	arr := create_array()
	create_array_using_make()
	//slice := make([]int, 5)
	//fmt.Printf("%p", &slice)
	reversed := reverse(arr[:])
	fmt.Println("reversed:", reversed)

	twoSumArr := [6]int{1, 2, 4, 7, 9, 20}
	result := twoSumSorted(twoSumArr[:], 27)
	fmt.Println(result)

	/* remove duplicates*/
	duplicateArr := []int{1, 1, 2, 2, 3}
	fmt.Println(removeDuplicates(duplicateArr))

	nums := []int{2, 1, 5, 1, 3, 2}
	k := 4
	fmt.Print(slidingWindow(nums, k))
}
