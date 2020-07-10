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
	shMem, err := shm.Create(4096)
	if err != nil {
		log.Fatal("Failed to create shared memory region")
	}

	pBase, err := shMem.Attach()
	if err != nil {
		log.Fatal("Failed to connect to shared memory region")
	}

	var i int32
	for i = 0; i < 100; i++ {
		index := uintptr(i)
		stride := uintptr(unsafe.Sizeof(int32(0)))

		addr := (*int32)(unsafe.Pointer(uintptr(pBase) + index*stride))
		*addr = i
	}

	fmt.Printf("Shared memory ID: %d\n", shMem.Id)
	fmt.Printf("Shared memory Size: %d\n", shMem.Size)

	scanner := bufio.NewReader(os.Stdin)
	fmt.Println("Press ENTER to Continue")
	scanner.ReadString('\n')

	shMem.Detach(pBase)
	shMem.Destroy()

	return
}
