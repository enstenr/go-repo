package main

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	mutex sync.Mutex
	total int
}

func (c *Counter) AddOne() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.total++
}

func main() {
	likes := &Counter{total: 0}

	for i := 0; i < 1000; i++ {
		go likes.AddOne()
	}
	time.Sleep(1 * time.Second)
	fmt.Println(likes.total)
	great := &Counter{total: 0}

	likes.AddOne()
	likes.AddOne()

	great.AddOne()

	fmt.Print(likes.total)
	fmt.Print(great.total)
}
