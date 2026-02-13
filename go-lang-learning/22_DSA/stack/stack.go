package stack

import (
	"fmt"
	"strings"
)

type Stack struct {
	Items []rune
}

func (s *Stack) Push(item rune) {
	s.Items = append(s.Items, item)
}

func (s *Stack) Pop() rune {
	if len(s.Items) == 0 {
		return 0
	}
	lastIdx := len(s.Items) - 1
	item := s.Items[lastIdx]
	s.Items = s.Items[:lastIdx]
	return item
}
func (s *Stack) IsEmpty() bool {
	return len(s.Items) == 0
}

func ReverseWithStack(input string) string {
	s := Stack{}
	for _, ch := range input {
		s.Push(ch)
	}
	var result strings.Builder
	for !s.IsEmpty() {
		result.WriteRune(s.Pop())
	}
	return result.String()
}
func main1() {
	test := "Hello World"
	fmt.Println(ReverseWithStack(test))
}
