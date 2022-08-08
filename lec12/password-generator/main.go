package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

const FILENAME string = "passwords.csv"

type entry struct {
	id       string
	username string
	url      string
	password string
}

func parseEntry(array []string) entry {
	parsedEntry := entry{id: array[0], username: array[1], url: array[2], password: array[3]}
	return parsedEntry
}

func check(err error) {
	if err != nil {
		// log.Println(err)
		log.Panicln(err)
	}
}

func readData() ([][]string, error) {
	f, err := os.Open(FILENAME)
	if err != nil {
		f, err = os.Create(FILENAME)
		check(err)
		log.Println("created file")
	}
	defer f.Close()
	r := csv.NewReader(f)
	arrays, err := r.ReadAll()
	if err != nil {
		return [][]string{}, err
	}
	return arrays, nil
}

func generatePassword(length int) string {
	generatedPassword := ""
	for i := 0; i < length; i++ {
		generatedPassword += getOneRandomLetter()
	}
	return generatedPassword
}

func getOneRandomLetter() string {
	var randomLetter string
	idx := rand.Intn(26)
	uppercase := rand.Intn(2) == 0
	if uppercase {
		randomLetter = string(idx + 65)
	} else {
		randomLetter = string(idx + 97)
	}
	return randomLetter
}

func main() {
	arrays, err := readData()
	rand.Seed(time.Now().UnixNano())
	// fmt.Printf("file: %v\n", f)
	check(err)
	entries := []entry{}
	for _, item := range arrays {
		entries = append(entries, parseEntry(item))
	}

	fmt.Printf("entries: %v\n", entries)
	randomLetter := generatePassword(8)
	fmt.Printf("randomLetter: %v\n", randomLetter)
}

// no, username, url, password
// 1, pwhb, github.com, helloworld
