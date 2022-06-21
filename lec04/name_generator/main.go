package main

import (
	"fmt"
	"math/rand"
	"time"
)

func getRandName() string {
	// generate random
	nowUnix := time.Now().UnixNano()
	seed := rand.NewSource(nowUnix)
	newRand := rand.New(seed)

	names := []string{"Aung", "Bo", "Chit", "Daung", "Ei", "Gone", "Htet"}
	total := len(names)
	randNum := newRand.Intn(total)
	// var randomName string = names[]
	// randomName := names[randNum]
	// return randomName
	return names[randNum]
}

func main() {
	num := 10
	var nameString string
	for i := 0; i < num; i++ {
		randName := getRandName()
		nameString = nameString + " " + randName
	}
	fmt.Printf("random name: %v\n", nameString)
	// fmt.Printf("total: %v\n", total)
}
