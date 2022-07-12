package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/pwhb/go-hangman/res"
)

func main() {
	fmt.Println("Welcome to hangman!")
	answer := getRandomWord()
	wordLength := len(answer)

	hangmanStatus := 0
	userWon := false

	// fmt.Printf("answer: %v\n", answer)

	// create an empty array with the length of the answer
	userAnswerArr := make([]string, wordLength)
	// fill the empty array with "_"
	for idx := range userAnswerArr {
		userAnswerArr[idx] = "_"
	}

	fmt.Printf("%v\n", res.Figures[hangmanStatus])

	for hangmanStatus < 6 {
		if checkUserInput(userAnswerArr, answer) {
			if isGameOver(userAnswerArr) {
				userWon = true
				break
			}
		} else {
			hangmanStatus++
		}
		fmt.Printf("\n%v\n", res.Figures[hangmanStatus])
	}

	if userWon {
		fmt.Println("You win!")
	} else {
		fmt.Printf("You lost.")
	}

	fmt.Printf("The answer is `%v`\n", answer)

}

func getRandomWord() string {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomIdx := random.Intn(21)
	return res.Words[randomIdx]
}

func checkUserInput(userAnswerArr []string, answer string) bool {
	result := false
	fmt.Println(strings.Join(userAnswerArr, " "))

	var userInput string

	fmt.Print("\nChoose an alphabet: ")

	fmt.Scanf("%v", &userInput)

	// ignore space
	userInput = strings.TrimSpace(userInput)

	for idx, val := range answer {
		if string(val) == userInput {
			result = true
			userAnswerArr[idx] = userInput
		}
	}
	return result
}

func isGameOver(userAnswerArr []string) bool {
	for _, val := range userAnswerArr {
		if val == "_" {
			return false
		}
	}
	return true
}
