package main

import "fmt"

type Stack struct {
	str string
}

func (s *Stack) reverse() string {
	runes := []rune(s.str)
	n := len(runes)
	j := n - 1
	for i := 0; i < n/2; i++ {
		runes[i], runes[j-i] = runes[j-i], runes[i]
	}
	return string(runes)
}
func main() {
	s := Stack{str: "Hello World"}
	fmt.Print(s.reverse())
}
