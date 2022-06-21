package main

import "fmt"

func main() {
	var i int16
	// for start, end, step
	for i = 0; i < 50; i += 5 {
		fmt.Printf("current num: %v\n", i)
	}
	// for i < 1000 {
	// 	fmt.Println(i)
	// 	i++
	// }
}
