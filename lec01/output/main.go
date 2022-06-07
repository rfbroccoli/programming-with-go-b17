package main

import "fmt"

func main() {
	welcomeText := "welcome to lecture"
	lecNum := 1
	lecturer := "sayar broccoli"
	// Print
	// fmt.Print(welcomeText)
	// fmt.Print(" ")
	// fmt.Print(lecNum)
	// fmt.Print(" by ")
	// fmt.Print(lecturer)
	// fmt.Print("!")
	// fmt.Print("\n")

	// Printf
	// fmt.Printf("%v %v by %v!\n", welcomeText, lecNum, lecturer)
	fmt.Printf("%v %v by %q!\n", welcomeText, lecNum, lecturer)

	// line break
	// fmt.Print("\n")

}
