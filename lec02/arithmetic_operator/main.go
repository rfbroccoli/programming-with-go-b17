package main

import "fmt"

func main() {
	a := 10
	b := 2
	c := a + b
	fmt.Printf("%v + %v = %v\n", a, b, c)

	c = a - b
	fmt.Printf("%v - %v = %v\n", a, b, c)

	c = a * b
	fmt.Printf("%v * %v = %v\n", a, b, c)

	c = a / b
	fmt.Printf("%v / %v = %v\n", a, b, c)

	c = a % b
	fmt.Printf("%v %% %v = %v\n", a, b, c)

	d := 10
	fmt.Printf("d = %v\n", d)

	// d = d + 5
	d += 5
	fmt.Printf("d = %v\n", d)
	// d = d {operator} {num}
	// d {operator}= {num}
	// d = d / 3
	d /= 3
	fmt.Printf("d = %v\n", d)

	// d++
	// fmt.Printf("d = %v\n", d)

	// d--
	// fmt.Printf("d = %v\n", d)
}
