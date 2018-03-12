package main

import (
	"os"
	"fmt"
)

func main(){
	fmt.Printf("Hello, %s", os.Args[1])
}