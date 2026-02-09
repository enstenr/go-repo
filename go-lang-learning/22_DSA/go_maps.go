package main

import (
	"fmt"
	"unsafe"
)

type hmap struct {
	count     int
	flags     uint8
	B         uint8 // log_2 of # of buckets (e.g., if B=3, buckets=8)
	noverflow uint16
	hash0     uint32         // hash seed
	buckets   unsafe.Pointer // THE MAIN ARRAY pointer
}
type bmap struct {
	tophash [8]uint8
	// After this, Go packs 8 Keys, then 8 Values
}

func main() {
	m := make(map[string]int)
	m["Rajesh"] = 40
	fmt.Println(m["Rajesh"])
	fmt.Println(unsafe.Sizeof(m))

	header := *(**hmap)(unsafe.Pointer(&m))
	fmt.Println(header)

	bucketPtr := header.buckets
	b := (*bmap)(bucketPtr)
	key0Add := uintptr(bucketPtr) + unsafe.Sizeof(b.tophash)
	rajeshHeader := (*uintptr)(unsafe.Pointer(key0Add))

	fmt.Printf("1. Map Header Address: %p\n", header)
	fmt.Printf("2. Bucket Address:     %p\n", bucketPtr)
	fmt.Printf("3. Key 0 Header Address: %p\n", rajeshHeader)

	strPtr := (*uintptr)(unsafe.Pointer(rajeshHeader))
	fmt.Printf("4. Actual String Characters ('Rajesh') are at: 0x%x\n", *strPtr)

	// Verify the value in the bucket
	// Offset = BucketStart + TopHash(8) + AllKeys(16*8)
	val0Addr := uintptr(bucketPtr) + 8 + (16 * 8)
	val0 := (*int)(unsafe.Pointer(val0Addr))
	fmt.Printf("5. Value at Slot 0:    %d\n", *val0)

}
