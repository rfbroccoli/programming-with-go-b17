package main

import "fmt"

func main() {
	num1 := 5
	num2 := 60
	// num3 := add(num1, num2)
	// fmt.Printf("%v + %v = %v\n", num1, num2, num3)
	fmt.Printf("%v add %v = %v\n", num1, num2, add(num1, num2))
	fmt.Printf("%v substract %v = %v\n", num1, num2, substract(num1, num2))
	fmt.Printf("%v multiply %v = %v\n", num1, num2, multiply(num1, num2))
	fmt.Printf("%v divide %v = %v\n", num1, num2, divide(num1, num2))
}

func add(n1 int, n2 int) int {
	return n1 + n2
}

func multiply(n1 int, n2 int) int {
	return n1 * n2
}

func divide(n1 int, n2 int) int {
	return n2 / n1
}

func substract(n1 int, n2 int) int {
	return n1 - n2
}
