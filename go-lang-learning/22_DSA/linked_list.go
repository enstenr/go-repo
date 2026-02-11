package main

import "fmt"

type rec struct {
	data     string
	previous *rec
	next     *rec
}

func main() {
	rec1 := rec{data: "Rajesh"}
	rec2 := rec{data: "Iswarya"}
	rec3 := rec{data: "Nishesha"}
	rec4 := rec{data: "Vishwesh"}
	rec1.next = &rec2
	rec2.previous = &rec1
	rec2.next = &rec3
	rec3.next = &rec4
	rec3.previous = &rec2
	rec4.previous = &rec3

	//travel through the list
	current := &rec1
	fmt.Printf("Rec1:%s %x\n", current.data, current)
	fmt.Printf("Address of Rec1: %p\n", &rec1)
	fmt.Printf("Address of Rec2: %p\n", &rec2)
	fmt.Printf("Pointer inside Rec1: %p\n", rec1.next)
}
