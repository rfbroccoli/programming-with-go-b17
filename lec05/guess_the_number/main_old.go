package main

import "fmt"

func main() {
	num := 34
	var guess int
	fmt.Print("The hidden number is in 0-99\n")

	for guess != num {
		fmt.Print("Type your guess: ")
		fmt.Scan(&guess)
		fmt.Printf("Your guess is %v\n", guess)
		// if correct >> print correct
		// if not correct >>

		// if less than 0 >> print guess cant be less than 0
		// if greater than 99 >> print guess cant greater than 99

		// if less than num >> print less than num
		// if greater than num >> print greater than num
	}
}
