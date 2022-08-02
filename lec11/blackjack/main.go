package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

var cards = [13]string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}

func main() {
	fmt.Println("Welcome to Blackjack!")
	rand.Seed(time.Now().UnixNano())
	// card array
	userHand := []string{}
	computerHand := []string{}
	stop := false
	userHand = append(userHand, dealACard())
	userHand = append(userHand, dealACard())

	computerHand = append(computerHand, dealACard())
	computerHand = append(computerHand, dealACard())
	// [firstCard secondCard]
	fmt.Printf("Player: %v, Score: %v\n", userHand, convertHandToScore(userHand))
	fmt.Printf("Dealer: [%v ?], Score: ?\n", computerHand[1])
	if convertHandToScore(userHand) == 21 && convertHandToScore(computerHand) == 21 {
		fmt.Println("Both got Blackjack. Draw!")
		return
	}

	if convertHandToScore(userHand) == 21 {
		fmt.Println("You got Blackjack. You won!")
		return
	}

	if convertHandToScore(computerHand) == 21 {
		fmt.Println("Dealer got Blackjack. You lost!")
		return
	}

	for !stop {
		var userInput string
		fmt.Print("[D]raw or [S]top? ")
		fmt.Scanf("%v", &userInput)
		fmt.Println(userInput)
		userInput = strings.TrimSpace(userInput)
		userInput = strings.ToUpper(userInput)
		switch userInput {
		case "D":
			userHand = append(userHand, dealACard())
			fmt.Printf("Player: %v, Score: %v\n", userHand, convertHandToScore(userHand))
			if convertHandToScore(computerHand) > 21 {
				stop = true
			}
		case "S":
			stop = true
		default:
			fmt.Println("Invalid Input. Please enter a valid input.")
		}
	}

	for convertHandToScore(computerHand) < 16 {
		computerHand = append(computerHand, dealACard())
		fmt.Printf("Dealer: %v, Score: %v\n", computerHand, convertHandToScore(computerHand))
	}

	if convertHandToScore(userHand) > 21 && convertHandToScore(computerHand) > 21 {
		fmt.Println("Both got busted. Draw!")
		return
	}

	if convertHandToScore(userHand) > 21 {
		fmt.Println("You got busted. You lost!")
		return
	}

	if convertHandToScore(computerHand) > 21 {
		fmt.Println("Dealer got busted. You won!")
		return
	}

	if convertHandToScore(userHand) == convertHandToScore(computerHand) {
		fmt.Println("Equal Score. Draw!")
		return
	}

	if convertHandToScore(userHand) > convertHandToScore(computerHand) {
		fmt.Println("You won!")
		return
	}

	if convertHandToScore(userHand) < convertHandToScore(computerHand) {
		fmt.Println("You lost!")
		return
	}
}

func dealACard() string {
	index := rand.Intn(13)
	randomCard := cards[index]
	return randomCard
}

func convertCardToScore(card string, firstTurn bool) int {
	switch card {
	case cards[0]:
		if firstTurn {
			return 11
		}
		return 1
	case cards[10]:
		{
			return 10
		}
	case cards[11]:
		return 10
	case cards[12]:
		return 10
	default:
		score, _ := strconv.Atoi(card) // score, _ := strconv.ParseInt(card, 10, 0)
		return score
	}
}

func convertHandToScore(hand []string) int {

	firstTurn := false
	if len(hand) == 2 {
		firstTurn = true
	}
	totalScore := 0
	for _, c := range hand {
		totalScore += convertCardToScore(c, firstTurn)
		// fmt.Printf("i: %v\n", i)
		// fmt.Printf("c: %v\n", c)
	}
	return totalScore
}
