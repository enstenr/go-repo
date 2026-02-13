package main

import (
	"fmt"
	"unsafe"
)

func main() {
	rawBytes := []byte("Iswarya Rajesh ")
	standardString := string(rawBytes)
	fmt.Println(&rawBytes[0], len(rawBytes))
	zerocopy := unsafe.String(&rawBytes[0], len(rawBytes))
	fmt.Print(standardString, zerocopy)
}
