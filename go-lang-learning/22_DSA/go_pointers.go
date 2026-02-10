package main

import (
	"fmt"
	"unsafe"
)

func main() {
	name := "Rajesh"
	fmt.Println(&name, name) //printing address of var name, name

	nameptr := &name                         //this stores address of var name
	*nameptr = "Iswarya"                     // this line is to udpate the value directly at address
	fmt.Println(&nameptr, nameptr, *nameptr) //printing address of var name, name

	ptr2 := &nameptr
	fmt.Println(&ptr2, ptr2, *ptr2, **ptr2)

	ptr3 := &ptr2
	fmt.Println(&ptr3, ptr3, *ptr3, **ptr3, ***ptr3)
	fmt.Println(**(**string)(unsafe.Pointer(ptr2)))

}
