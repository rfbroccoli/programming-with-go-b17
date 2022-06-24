package main

import (
	"fmt"
)

func main() {
	var i string

	fmt.Print("Type in your name: ")
	fmt.Scanf("My name is %v", &i)
	fmt.Println("Your name is:", i)
}
