package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unsafe"

	"github.com/ghetzel/shmtool/shm"
)

func main() {
	shMem, err := shm.Create(800)
	if err != nil {
		log.Fatal("Failed to create shared memory region")
	}

	pBase, err := shMem.Attach()
	if err != nil {
		log.Fatal("Failed to connect to shared memory region")
	}

	var i float64
	for i = 0; i < 100; i++ {
		index := uintptr(i)
		stride := uintptr(unsafe.Sizeof(float64(0)))

		addr := (*float64)(unsafe.Pointer(uintptr(pBase) + index*stride))
		*addr = i
	}

	fmt.Printf("Shared memory ID: %d\n", shMem.Id)
	fmt.Printf("Shared memory Size: %d\n", shMem.Size)

	scanner := bufio.NewReader(os.Stdin)
	fmt.Println("Press ENTER to Continue")
	scanner.ReadString('\n')

	for j := 0; j < 100; j++ {
		pAddr := (*[100]float64)(pBase)
		arr := *pAddr

		fmt.Printf("NUM: %f\n", arr[j])
	}

	shMem.Detach(pBase)
	shMem.Destroy()

	return
}
