package main

import (
	"fmt"
)

func main() {
	fmt.Println(" Fibonacci number Series Program in Go Language")
	var totalNumber int
	fmt.Scan(&totalNumber)

	num1 := 0
	num2 := 1
	//First two numbers are printed initially
	fmt.Println(num1)
	fmt.Println(num2)
	sum := num1 + num2

	for startValue:=1;startValue < totalNumber;startValue++ {
		fmt.Println(sum)
		num1 = num2
		num2 = sum
		sum = num1 + num2

	}

}
