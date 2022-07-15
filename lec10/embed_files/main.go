package main

import (
	"embed"
	"fmt"
	"log"
)

//go:embed hello.txt
var f embed.FS

func main() {
	bytes, err := f.ReadFile("hello.txt")
	if err != nil {
		log.Println(err)
	}
	str := string(bytes)
	fmt.Printf("string: %v\n", str)
}
