package main

import (
	"fmt"
)

func main() {
	// sayWelcome()
	sayHelloTo("soe thiha swe")
	sayHelloTo("jde 16-0202")
	sayHelloTo("hein htoo")
}

func sayHelloTo(name string) {
	fmt.Printf("hello, %v!\n", name)
}

func sayHello() {
	fmt.Println("hello")
	fmt.Println("welcome")
	fmt.Println("to")
	fmt.Println("programming")
	fmt.Println("with")
	fmt.Println("go")
	fmt.Println("lecture")
	fmt.Println("03")
	fmt.Println("")
}
