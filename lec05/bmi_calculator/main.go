package main

import "fmt"

// Underweight (Severe thinness) 	< 16.0 	< 0.64
// Underweight (Moderate thinness) 	16.0 – 16.9 	0.64 – 0.67
// Underweight (Mild thinness) 	17.0 – 18.4 	0.68 – 0.73
// Normal range 	18.5 – 24.9 	0.74 – 0.99
// Overweight (Pre-obese) 	25.0 – 29.9 	1.00 – 1.19
// Obese (Class I) 	30.0 – 34.9 	1.20 – 1.39
// Obese (Class II) 	35.0 – 39.9 	1.40 – 1.59
// Obese (Class III) 	≥ 40.0 	≥ 1.60

func main() {

	// mass := 190.0
	// height := 1.80

	var mass float32
	fmt.Print("Type your weight in kg: ")
	fmt.Scan(&mass)
	var height float32
	fmt.Print("Type your height in m: ")
	fmt.Scan(&height)

	bmi := mass / (height * height)

	fmt.Printf("Your BMI is %v\n", bmi)
	fmt.Print("You're ")
	if bmi < 16.0 {
		fmt.Println("Underweight (Severe thinness)")
	} else if bmi < 17 {
		fmt.Println("Underweight (Moderate thinness)")
	} else if bmi < 18.5 {
		fmt.Println("Underweight (Mild thinness)")
	} else if bmi < 25 {
		fmt.Println("Normal range")
	} else if bmi < 30 {
		fmt.Println("Overweight (Pre-obese)")
	} else if bmi < 35 {
		fmt.Println("Obese (Class I)")
	} else if bmi < 40 {
		fmt.Println("Obese (Class II)")
	} else {
		fmt.Println("Obese (Class III)")
	}
}
