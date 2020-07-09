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
	//
	// Slice is a struct, composed of:
	// - len (INT, tells the length of a slice)
	// - cap (INT, tells the capacity of a slice)
	// - arr ([]ANY, underlying array that holds actual data)
	//
	// Therefore, &mat IS NOT ADDRESS OF first element in mat
	// &mat[0][0] is the address of the first element !

	for i := 0; i < 100; i++ {
		// Use uintptr for pointer arithmetic
		stride := uintptr(unsafe.Sizeof(int(0)))
		index := uintptr(i)

		num := *(*int)(unsafe.Pointer(uintptr(pMat) + index*stride))
		// get memory address by uintptr operation,
		// then cast it into unsafe pointer,
		// then cast it into *int,
		// then dereference it.

		fmt.Printf("%d\n", num)
	}

	return
}
