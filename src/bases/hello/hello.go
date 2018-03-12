package main

import (
	"fmt"
	"os"
)

func main() {
	var name string = "World"

	//len() === lenght
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	//fmt === format
	fmt.Printf("Hello, %s", name)
}
