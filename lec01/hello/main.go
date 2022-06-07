package main

import (
	"fmt"
)

func main() {
	// this line prints "welcome to lecture 01!"
	const welcome = "welcome to lecture"

	var lectureNum = 01
	fmt.Println(welcome, lectureNum, "!")

	lectureNum = 02
	fmt.Println(welcome, lectureNum, "!")

	fmt.Println(23)
	fmt.Println(false)

	var snake_case = "Snake Case"
	fmt.Println(snake_case)

	var camelCase = "Camel Case"
	fmt.Println(camelCase)

	var PascalCase = "Pascal Case"
	fmt.Println(PascalCase)
}
