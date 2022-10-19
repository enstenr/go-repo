package main

import ("fmt"
"time")

func main() {
	nums:=[]int64{1,2,3,4,5,6,7,8,9,10,11}
	for i, num := range nums {
        
            fmt.Println("index:", i,num)
    
    }
	fmt.Println(time.Now())
}