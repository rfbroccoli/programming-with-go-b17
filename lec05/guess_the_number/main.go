package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	nowUnix := time.Now().UnixNano()
	seed := rand.NewSource(nowUnix)
	newRand := rand.New(seed)

	num := newRand.Intn(100)
	var guess int

	// TODO
	// DONE Copy the codes from lec04 to generate the variable num randomly each time
	// DONE Add game logic so the player is given a valid response so they can make a better guess
	// DONE Add some logic to remind the player that the number is between 0 and 99, both ends included
	// DONE Add tries and lives
	tries := 0
	lives := 7
	fmt.Println("The hidden number is in 0-99")
	for true {
		fmt.Print("Type in your guess: ")
		fmt.Scan(&guess)
		tries++

		if guess == num {
			fmt.Printf("The number you guessed, %v, is correct!\nYou got the right answer in %v tries.\n", guess, tries)
			break
		}
		// lives = lives - 1
		lives--
		if num < 0 {
			fmt.Print("Your guess can't be less than 0.\n")
		} else if num > 99 {
			fmt.Print("Your guess can't be greater than 99.\n")
		} else if num > guess {
			fmt.Printf("The hidden number is greater than your guess, %v.\n", guess)
		} else {
			fmt.Printf("The hidden number is less than your guess, %v.\n", guess)
		}
		fmt.Printf("You have %v chances left.\n", lives)
		if lives == 0 {
			fmt.Printf("The hidden number is %v.\n", num)
			break
		}
	}

}
