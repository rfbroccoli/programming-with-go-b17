package main

import "fmt"

func main() {
	a := 10
	b := 14

	fmt.Printf("a = %v\n", a)
	fmt.Printf("b = %v\n", b)

	fmt.Printf("a > b = %v\n", a > b)
	fmt.Printf("a < b = %v\n", a < b)
	fmt.Printf("a == b = %v\n", a == b)
	fmt.Printf("a != b = %v\n", a != b)
	fmt.Printf("a >= b = %v\n", a >= b)
	fmt.Printf("a <= b = %v\n", a <= b)

	str1 := "hello"
	str2 := "world"
	fmt.Printf("str1 = %v\n", str1)
	fmt.Printf("str2 = %v\n", str2)
	fmt.Printf("str1 == str2 = %v\n", str1 == str2)
	fmt.Printf("str1 != str2 = %v\n", str1 != str2)
}
