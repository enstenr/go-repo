package main

import "fmt"

func fibo(n int) ([]int, int) { //return slice and sum of all numbers in slice
	count := 2
	//fib := []int{0, 1}
	//fib := make([]int, count) // initializing to reduce the cost alloc/op
	fib := make([]int, count, n+2) // setting capacity
	a := 0
	b := 1
	c := 0
	sum := 1
	//if n = 0 return array of 0 and sum 0
	// if n = 1 then return 1 and sum 1
	if n == 0 {
		return []int{0}, 0
	}
	if n == 1 {
		return []int{0, 1}, 1
	}
	for i := 0; i <= n-2; i++ {

		c = a + b
		sum += c

		fib = append(fib, c)
		a = b
		b = c
		count++
	}

	//return sum of the array fib
	return fib, sum
}
func main() {
	for i := 0; i < 10; i++ {
		fib, sum := fibo(i)
		fmt.Println(fib, sum)
	}

}
