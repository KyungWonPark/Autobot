package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Foo is: ", os.Getenv("FOO"))
	fmt.Println("Bar is: ", os.Getenv("BAR"))

	return
}
