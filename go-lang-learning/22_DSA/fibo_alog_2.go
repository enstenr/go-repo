package main

import "fmt"

func fibo_2(n int, cache map[int]int) int { //return slice and sum of all numbers in slice

	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	value, ok := cache[n]
	if ok {
		return value
	} else {
		cache[n] = fibo_2(n-1, cache) + fibo_2(n-2, cache)
	}

	return cache[n]
}
func main() {
	cache := make(map[int]int)
	total := 0
	for i := 0; i < 10; i++ {

		fibonum := fibo_2(i, cache)

		fmt.Println(fibonum)
		total += fibonum

	}
	fmt.Println(total)

}
