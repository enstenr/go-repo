package main

import(
	"fmt"
)
func main(){
	greeting := "hello world!"
	name:="Rajesh"
	sayGreeting(greeting,&name)
	fmt.Println(name)
}

func sayGreeting(greeting string, name *string) {
	fmt.Println(greeting,*name)
	*name="Vishwesh"
	fmt.Println(*name)
}