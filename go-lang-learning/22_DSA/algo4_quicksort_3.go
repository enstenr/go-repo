package main

import "fmt"

func quicksort_4(arr []int) []int {

	if len(arr) < 2 {
		return arr
	}
	arr = presort(arr)
	pivot := arr[len(arr)-1]
	i := -1

	for j := 0; j < len(arr)-1; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]

		}
	}
	arr[len(arr)-1], arr[i+1] = arr[i+1], arr[len(arr)-1]
	quicksort_4(arr[:i+1])
	quicksort_4(arr[i+2:])
	return arr
}
func presort(arr []int) []int {
	//fmt.Println("pre ", arr)
	mid := len(arr) / 2
	last := len(arr) - 1
	if arr[0] > arr[mid] {
		arr[0], arr[mid] = arr[mid], arr[0]
	}
	if arr[0] > arr[last] {
		arr[0], arr[last] = arr[last], arr[0]
	}
	if arr[mid] > arr[last] {
		arr[mid], arr[last] = arr[last], arr[mid]
	}
	arr[mid], arr[last] = arr[last], arr[mid]
	//fmt.Println("post ", arr)
	return arr
}
func main() {
	arr := []int{11, 22, 3, 34, 5, 33, 12, 44, 54, 44, 99, 12, 15, 17, 81}

	fmt.Println(quicksort_4(arr))
}
