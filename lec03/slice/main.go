package main

import "fmt"

func main() {
	// var numArray = [6]int{}
	numSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("numSlice = %v\n", numSlice)
	numSlice[0] = 35
	fmt.Printf("numSlice[0] = %v\n", numSlice[0])
	fmt.Printf("numSlice[1] = %v\n", numSlice[1])
	fmt.Printf("numSlice[2] = %v\n", numSlice[2])
	fmt.Printf("numSlice[3] = %v\n", numSlice[3])
	fmt.Printf("numSlice[4] = %v\n", numSlice[4])
	// fmt.Printf("numSlice[5] = %v\n", numSlice[5])

	fmt.Printf("numSlice = %v\n", numSlice)
	fmt.Printf("len(numSlice) = %v\n", len(numSlice))
	// fmt.Printf("numArray[6] = %v\n", numArray[6])
	// stringArray := [3]string{"jde 16-0202", "soe thiha swe", "hein htoo"}

	// fmt.Printf("stringArray = %v\n", stringArray)
	// fmt.Printf("stringArray[2] = %v\n", stringArray[2])
	// stringArray[2] = "someone"
	// fmt.Printf("stringArray[2] = %v\n", stringArray[2])
}
