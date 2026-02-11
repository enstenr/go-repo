package main

import "fmt"

func quicksort_2(arr []int) []int {

	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {

			if arr[j] < arr[i] {
				arr[i], arr[j] = arr[j], arr[i]
			}

		}
	}
	return arr
}
func main() {
	fmt.Println(quicksort_2([]int{11, 22, 3, 34, 5, 33, 12, 44, 54}))
}
