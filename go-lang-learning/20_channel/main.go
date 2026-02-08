package main

import (
	"fmt"
	"sync"
)

func main() {

	pings_func := make(chan int)
	done := make(chan bool)
	var wg sync.WaitGroup
	total := 0

	go func() {
		for value := range pings_func {

			total = total + value

		}
		done <- true

	}()

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			pings_func <- 1
		}()
		//time.Sleep(time.Second)
	}
	wg.Wait()
	close(pings_func)
	<-done
	fmt.Println(total)

}
