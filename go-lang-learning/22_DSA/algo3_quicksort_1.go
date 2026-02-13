package main

import "fmt"

func quicksort(a []int) []int {
	//fmt.Println(a)
	sorted := []int{}
	if len(a) < 2 {
		//fmt.Println(" length less than 2 ", a)
		return a
	}
	pivot := a[0]

	var left, right []int
	for i := 1; i < len(a); i++ {
		if a[i] < pivot {
			left = append(left, a[i])
		} else {
			right = append(right, a[i])
		}
	}
	//fmt.Println("pivot ", pivot)
	//fmt.Println(" left ", left)
	//fmt.Println(" right ", right)
	sorted = append(quicksort(left), pivot)
	//fmt.Println(" This will be entered after finishing recursive loop above ", sorted)
	sorted = append(sorted, quicksort(right)...)
	//quicksort(right))
	return sorted
}

func main() {
	sorted := quicksort([]int{11, 22, 3, 34, 5})
	fmt.Println(sorted)
}
