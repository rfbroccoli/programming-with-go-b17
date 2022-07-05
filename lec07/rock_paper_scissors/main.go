package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// var choices [3]string = [3]string{"ROCK", "PAPER", "SCISSORS"}
// type choicesStruct struct {
// 	R string
// 	P string
// 	S string
// }

// var choices = choicesStruct{R: "ROCK", P: "PAPER", S: "SCISSORS"}

var choices = struct {
	R string
	P string
	S string
}{"ROCK", "PAPER", "SCISSORS"}

func getComputerChoice() string {

	nowUnix := time.Now().UnixNano()
	seed := rand.NewSource(nowUnix)
	newRand := rand.New(seed)
	computerChoiceIndex := newRand.Intn(3)
	// computerChoice := choices[computerChoiceIndex]

	switch computerChoiceIndex {
	case 0:
		return choices.R
	case 1:
		return choices.P
	case 2:
		return choices.S
	default:
		return ""
	}
}

func getUserChoice() string {

	var userChoice string
	fmt.Print("Choose a move from [R]ock, [P]aper, [S]cissors: ")
	fmt.Scanf("%s", &userChoice)

	// fmt.Printf("userChoice: %v\n", userChoice)
	userChoice = strings.ToUpper(userChoice)
	// fmt.Printf("userChoice: %v\n", userChoice)
	switch userChoice {
	case "R":
		userChoice = choices.R
	case "P":
		userChoice = choices.P
	case "S":
		userChoice = choices.S
	default:
		userChoice = ""
		fmt.Println("Invalid move.")
	}
	return userChoice
}

func main() {
	userChoice := getUserChoice()

	if userChoice == "" {
		return
	}

	computerChoice := getComputerChoice()

	fmt.Printf("you: %v vs computer: %v\n", userChoice, computerChoice)
	if userChoice == computerChoice {
		fmt.Println("Draw.")
	} else if (userChoice == choices.R && computerChoice == choices.S) || (userChoice == choices.P && computerChoice == choices.R) || (userChoice == choices.S && computerChoice == choices.P) {
		fmt.Println("You win!")
	} else {
		fmt.Println("You lose. :(")
	}

	// TODO
	// Add logic to response if the player wins or lose or draw

	// HINT
	// when win ?
	// 1) player rock && computer scissors
	// 2)        paper && computer rock
	// 3)        scissors && computer paper
	// if (userChoice == choices[0] && computerChoice == choices[2]) || case2 || case3

}
