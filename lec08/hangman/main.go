package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/pwhb/go-hangman/res"
)

func main() {
	// fmt.Println(len(res.Figures))
	// fmt.Println(len(res.Words))
	fmt.Println("Welcome to hangman!")
	randomWord := getRandomWord()
	wordLength := len(randomWord)

	hangmanStatus := 0

	fmt.Printf("randomWord: %v\n", randomWord)
	// fmt.Printf("wordLength: %v\n", wordLength)
	answer := make([]string, wordLength)

	// fmt.Printf("answer: %v\n", answer)

	// fmt.Printf("len(answer): %v\n", len(answer))
	fmt.Printf("%v\n", res.Figures[hangmanStatus])

	for hangmanStatus < 6 {
		userInput := getUserInput(answer)
		fmt.Printf("userInput: %v", userInput)
		// TODO check if userInput is in randomWord
		// if true --> update answer
		// else
		hangmanStatus++
		fmt.Printf("\n%v\n", res.Figures[hangmanStatus])
	}
	fmt.Printf("You lost.")

}

func getRandomWord() string {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomIdx := random.Intn(21)
	return res.Words[randomIdx]
}

func getUserInput(answer []string) string {
	for _, val := range answer {
		// fmt.Printf("idx: %v\n", idx)
		if val == "" {

			fmt.Printf("_ ")
		} else {
			fmt.Printf("%v ", val)
		}
	}
	var userInput string
	fmt.Print("\nChoose an alphabet: ")
	fmt.Scanf("%v", &userInput)
	return userInput
}
