package main

import "fmt"

func main() {
	// mon, tue, wed, thu, fri, sat, sun
	today := 35
	switch today {
	case 0:
		fmt.Println("today is monday")

	case 1:
		fmt.Println("today is tuesday")

	case 2:
		fmt.Println("today is wednesday")

	case 3:
		fmt.Println("today is thursday")

	case 4:
		fmt.Println("today is friday")

	case 5:
		fmt.Println("today is saturday")

	case 6:
		fmt.Println("today is sunday")
	default:
		fmt.Println(today)
	}
}
