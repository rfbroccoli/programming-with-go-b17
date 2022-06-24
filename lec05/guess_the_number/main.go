package main

import "fmt"

func main() {
	num := 34
	var guess int

	// TODO
	// Copy the codes from lec04 to generate the variable num randomly each time
	// Add game logic so the player is given a valid response so they can make a better guess
	// Add some logic to remind the player that the number is between 0 and 99, both ends included

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
