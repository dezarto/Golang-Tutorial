package main

import "fmt"

func main() {
	/* var studentName string = "John Doe"
	var grade float32 = 77
	var isPassed bool = true */

	/* var (
		studentName string  = "John Doe"
		grade       float32 = 77
		isPassed    bool    = true
	) */

	/* studentName := "John Doe"
	grade := 77
	isPassed := true
	*/

	//var studentName, grade, isPassed = "John Doe", 77, true

	studentName, grade, isPassed := "John Doe", 77, true

	fmt.Printf("Student name: %s, Grade: %f, Is passed: %t", studentName, grade, isPassed)
}
