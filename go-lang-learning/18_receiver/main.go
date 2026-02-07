package main

import "fmt"

type Counter struct {
	total int
}

func (c *Counter) AddOne() {
	c.total++
}

func main() {
	likes := &Counter{total: 0}
	great := &Counter{total: 0}

	likes.AddOne()
	likes.AddOne()

	great.AddOne()

	fmt.Print(likes.total)
	fmt.Print(great.total)
}
