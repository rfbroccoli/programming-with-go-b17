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

	fmt.Printf("nowUnix: %v\n", nowUnix)
	// randNum := rand.Int()
	// randNum := newRand.Int()
	randNum := newRand.Intn(100)
	fmt.Println(randNum)
}
