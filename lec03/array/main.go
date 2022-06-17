package main

import "fmt"

func main() {
	// var numArray = [6]int{}
	numArray := [6]int{}
	fmt.Printf("numArray = %v\n", numArray)
	numArray[0] = 35
	fmt.Printf("numArray[0] = %v\n", numArray[0])
	fmt.Printf("numArray[1] = %v\n", numArray[1])
	fmt.Printf("numArray[2] = %v\n", numArray[2])
	fmt.Printf("numArray[3] = %v\n", numArray[3])
	fmt.Printf("numArray[4] = %v\n", numArray[4])
	fmt.Printf("numArray[5] = %v\n", numArray[5])

	fmt.Printf("numArray = %v\n", numArray)
	fmt.Printf("len(numArray) = %v\n", len(numArray))
	// fmt.Printf("numArray[6] = %v\n", numArray[6])
	stringArray := [3]string{"jde 16-0202", "soe thiha swe", "hein htoo"}

	fmt.Printf("stringArray = %v\n", stringArray)
	fmt.Printf("stringArray[2] = %v\n", stringArray[2])
	stringArray[2] = "someone"
	fmt.Printf("stringArray[2] = %v\n", stringArray[2])
}
