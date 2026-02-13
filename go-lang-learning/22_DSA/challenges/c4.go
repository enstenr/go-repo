package main

import (
	"dsa/stack"
	"fmt"
	"math"
)

func DigPow(n int, p int) int {
	s := &stack.Stack{}
	sum := 0

	quotient := n
	for quotient > 0 {
		s.Push(rune(quotient % 10))
		quotient = (quotient / 10)

	}
	fmt.Println(s)
	for !s.IsEmpty() {
		number := s.Pop()
		sum += int(math.Pow(float64(number), float64(p)))
		p++
	}
	fmt.Println(sum)
	if sum%n == 0 {
		return sum / n
	}
	return -1

}

func main() {
	//fmt.Println(DigPow(89, 1))
	fmt.Println(DigPow(46288, 3))
	//fmt.Println(659 % 10)
}
