package main

import (
	"fmt"
	"time"
)

func main() {
	// mon, tue, wed, thu, fri, sat, sun
	today := time.Now().Weekday()
	// today := time.Saturday
	fmt.Printf("today: %v\n", today)
	switch today {
	case time.Monday:
		fmt.Println("there's a class tommorow")

	case time.Tuesday:
		fmt.Println("there's a class today")

	case time.Wednesday:
		fmt.Println("there's a class coming friday")

	case time.Thursday:
		fmt.Println("there's a class tommorow")

	case time.Friday:
		fmt.Println("there's a class today")

	case time.Saturday:
		fmt.Println("there's a class coming tuesday")

	case time.Sunday:
		fmt.Println("there's a class coming tuesday")
	}
}
