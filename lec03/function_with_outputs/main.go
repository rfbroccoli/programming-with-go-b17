package main

import "fmt"

func main() {
	fmt.Println("function with output")
	output := funcWithOutput()
	fmt.Println(output)
}

func funcWithOutput() string {
	// fmt.Println("print something")
	// fmt.Println("print something")
	return "this is an output"
}
