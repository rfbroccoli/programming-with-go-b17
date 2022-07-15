package main

import (
	"fmt"
	"os"
)

func main() {
	bytes, _ := os.ReadFile("hello.txt")
	fmt.Printf("bytes: %v\n", bytes)
	fmt.Printf("string(bytes): %v\n", string(bytes))
	// fmt.Printf("err: %v\n", err)

}
