package res

import (
	"embed"
	"log"
	"math/rand"
	"strings"
	"time"
)

//go:embed words.txt
var f embed.FS

func getWords() []string {
	bytes, err := f.ReadFile("words.txt")
	if err != nil {
		log.Printf("error: %v", err)
	}
	return strings.Split(string(bytes), "\n")
}

func GetRandomWord() string {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	words := getWords()
	randomIdx := random.Intn(len(words))
	return words[randomIdx]
}
