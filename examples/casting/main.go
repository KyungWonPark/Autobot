package main

import (
	"fmt"
	"math/rand"
	"unsafe"
)

func main() {
	bytes := []byte{104, 101, 108, 108, 111}

	p := unsafe.Pointer(&bytes)
	str := *(*string)(p) //cast it to a string pointer and assign the value of this pointer
	fmt.Println(str)     //prints "hello"

	var mat [10][10]int
	for i := range mat {
		for j := range mat[i] {
			mat[i][j] = rand.Intn(100)
		}
	}

	pMat := unsafe.Pointer(&mat[0][0])
	// Remember! slice IS A CONTAINER TYPE !
	// Be aware of difference between Go array and slice
	// Slice is a container type, that is wrapping Go array
	//
	// Slice is a struct, composed of:
	// - len (INT, tells the length of a slice)
	// - cap (INT, tells the capacity of a slice)
	// - arr ([]ANY, underlying array that holds actual data)
	//
	// Therefore, &mat IS NOT AN ADDRESS OF first element in mat
	// &mat[0][0] is the address of the first element !

	// One way
	for i := 0; i < 100; i++ {
		arr := *(*[100]int)(pMat)
		// Note that we can dereference pMat as an ARRAY (which is just like C array, continuous static list of elements)
		// but not as SLICE ! *(*[]int)(pMat) will result in error due to the reason explained above.
		// If you can't understand, draw memory diagram and think again.

		fmt.Printf("%d\n", arr[i])
	}

	// Another way with manual pointer arithmetic
	for i := 0; i < 100; i++ {
		// uintptr is just uint, representing memory address
		index := uintptr(i)                      // which is i
		stride := uintptr(unsafe.Sizeof(int(0))) // which is 4 (bytes)

		num := *(*int)(unsafe.Pointer(uintptr(pMat) + index*stride))
		// Convert pMat unsafe pointer to uintptr for pointer arithmetic,
		// start from address: pMat, go (index * stride) bytes,
		// convert that address into unsafe pointer,
		// then cast it as an pointer pointing to integer data,
		// then dereference it.

		fmt.Printf("%d\n", num)
	}

	// Caution:
	// Golang has garbage collection.
	// If an array that your unsafe pointer is pointing gets GC'd,
	// you'll have segmentation fault.
	// So be cautious of variable scope
	// Don't use pointer casting outside of data's lifespan.
	// Go GCs variables when they are out-of-reach

	// Read:
	//
	// https://blog.gopheracademy.com/advent-2017/unsafe-pointer-and-system-calls/
	// https://blog.gopheracademy.com/advent-2017/unsafe-pointer-and-system-calls/

	return
}
