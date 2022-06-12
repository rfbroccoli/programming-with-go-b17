package main

import "fmt"

func main() {
	a := true
	b := false

	fmt.Printf("a = %v\n", a)
	fmt.Printf("b = %v\n", b)
	// NOT
	fmt.Printf("!a = %v\n", !a)
	fmt.Printf("!b = %v\n", !b)
	// OR ||
	fmt.Printf("a || b = %v\n", a || b)
	fmt.Printf("true || true = %v\n", true || true)
	fmt.Printf("true || false = %v\n", true || false)
	fmt.Printf("false || true = %v\n", false || true)
	fmt.Printf("false || false = %v\n", false || false)
	// AND &&
	fmt.Printf("a && b = %v\n", a && b)
	fmt.Printf("true && true = %v\n", true && true)
	fmt.Printf("true && false = %v\n", true && false)
	fmt.Printf("false && true = %v\n", false && true)
	fmt.Printf("false && false = %v\n", false && false)
}
