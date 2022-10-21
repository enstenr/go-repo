package main

import (
	"calculate"
	"fmt"
)

func main() {
	num1 := 0
	num2 := 0

	fmt.Scan(&num1)
	fmt.Scan(&num2)

	
	returnValue:=calculate.Calculate(num1, num2, "ADD")
	fmt.Println(returnValue)
}
