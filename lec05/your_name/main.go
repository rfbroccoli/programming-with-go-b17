package main

import (
	"fmt"
)

func main() {
	// var i, j string
	var i string
	var j string

	for (i != "your") || (j != "name") {

		fmt.Print("Type your name: ")
		fmt.Scanln(&i, &j)
	}
	fmt.Println("Bingo!")
}
