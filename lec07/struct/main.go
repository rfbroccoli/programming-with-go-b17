package main

import "fmt"

type student struct {
	name       string
	student_id int
	classes    []string
	year       string
}

func main() {
	// var studentObj student = student{}
	// studentObj.name = "Hein Htoo"
	// studentObj.student_id = 1234

	// studentObj.classes = []string{"CS", "Maths", "English"}
	// studentObj.year = "Third Year"
	// studentObj := student{name: "Hein Htoo", student_id: 1234, year: "Third Year", classes: []string{"CS", "Maths", "English"}}
	studentObj := student{"Hein Htoo", 1234, []string{"CS", "Maths", "English"}, "Third Year"}

	fmt.Printf("studentObj: %v \n", studentObj)
	fmt.Printf("studentObj.classes[0]: %v \n", studentObj.classes[0])
	// fmt.Printf("studentObj.nastudentObj: %v \n", studentObj.nastudentObj)
	// fmt.Printf("studentObj.job: %v \n", studentObj.job)
}
