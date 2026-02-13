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
	fmt.Println("Size of map ", unsafe.Sizeof(m))

	fmt.Printf("Address of Map %p \n", &m)
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

	// Scan all 8 possible slots in the bucket
	fmt.Println("--- Raw Bucket Dump (First 128 bytes) ---")
	for i := 0; i < 16; i++ { // Look at 16 chunks of 8-bytes
		ptr := uintptr(bucketPtr) + uintptr(i*8)
		val := (*uintptr)(unsafe.Pointer(ptr))
		fmt.Printf("Offset %d: 0x%x\n", i*8, *val)
	}

}
