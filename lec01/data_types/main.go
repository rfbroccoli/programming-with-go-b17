package main

import "fmt"

func main() {
	// var undeclared = 234
	// undeclared = "hello"
	// fmt.Print(undeclared)

	// TEXT STRING
	var str string = "this is a string"
	// NUMBER INT FLOAT
	var num int8 = 23
	var notInt float32 = 34.643
	// CONDITION BOOL
	var condition bool = true

	// var newString string = "this is a new string"
	newString := "this is a new string"
	// var newNum int = 45
	newNum := 45
	// var newCondition bool = false
	newCondition := false
	// var newFloat float64 = 45
	newFloat := 45.634

	// fmt.Println(str + str)
	// fmt.Println(num + num)
	// fmt.Println(condition)

	fmt.Println(str)
	fmt.Println(num)
	fmt.Println(notInt)
	fmt.Println(condition)

	fmt.Println(newString)
	fmt.Println(newNum)
	fmt.Println(newFloat)
	fmt.Println(newCondition)
}
