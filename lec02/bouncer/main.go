package main

import "fmt"

func main() {
	age := 23

	if age < 18 {
		fmt.Println("underage")
	} else {
		fmt.Println("allowed to drink")
	}
}
