package main

import "fmt"

func main() {
	var num int64
	fmt.Println("Enter the Number ")
	fmt.Scanln(&num)

	factorialValue := fact(num)
	fmt.Println()
	fmt.Printf(" Factorial of the Number is %d",factorialValue)
	fmt.Println()
}
func fact(num int64) int64 {
	if num == 0 {
		return 1
	}
	return num * fact(num-1)

}
