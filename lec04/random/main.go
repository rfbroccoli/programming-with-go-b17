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
	randNum := newRand.Intn(100)

	fmt.Printf("nowUnix: %v\n", nowUnix)
	// randNum := rand.Int()
	// randNum := newRand.Int()
	fmt.Println(randNum)
}
