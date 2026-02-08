package main

import (
	"fmt"
	"unsafe"
)

func create_array() {
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
}
func create_array_using_make() {
	s := make([]int, 0, 2)
	fmt.Printf("Initial addr = %p, len = %d , cap = %d \n", s, len(s), cap(s))

	for i := 0; i < 5; i++ {
		s = append(s, i)
		fmt.Printf("Initial addr = %p, len = %d , cap = %d \n", s, len(s), cap(s))
	}
}
func main() {
	//create_array()
	create_array_using_make()
	//slice := make([]int, 5)
	//fmt.Printf("%p", &slice)
}
