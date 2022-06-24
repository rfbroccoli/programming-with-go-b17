package main

import "fmt"

func main() {
	num := 34
	var guess int

	fmt.Println("The hidden number is in 0-99")
	for guess != num {
		fmt.Print("Type in your guess: ")
		fmt.Scan(&guess)
	}
	// fmt.Printf("The number you guessed, %v, is correct!\n", guess)
	// fmt.Printf("The hidden number is less than your guess, %v.\n", guess)
	// fmt.Printf("The hidden number is greater than your guess, %v.\n", guess)
	// fmt.Print("Your guess can't be greater than 99.\n")
	// fmt.Print("Your guess can't be less than 0.\n")

}
