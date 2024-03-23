package main

import "fmt"

func main() {

	var students, secondGroup, thirdGroup, fourthGroup, fail int
	var grade, totalGrade float64

	fmt.Scanln(&students)

	for i := 0; i < students; i++ {
		fmt.Scanln(&grade)
		totalGrade += grade

		if grade <= 2.99 {
			fail++
		} else if grade >= 3.00 && grade <= 3.99 {
			secondGroup++
		} else if grade >= 4.00 && grade <= 4.99 {
			thirdGroup++
		} else if grade >= 5.00 {
			fourthGroup++
		}
	}

	totalStudents := float64(students)
	averageGrade := totalGrade / totalStudents
	failPercent := float64(fail) / totalStudents * 100

	fmt.Printf("Top students: %.2f%%\n", float64(fourthGroup)/totalStudents*100)
	fmt.Printf("Between 4.00 and 4.99: %.2f%%\n", float64(thirdGroup)/totalStudents*100)
	fmt.Printf("Between 3.00 and 3.99: %.2f%%\n", float64(secondGroup)/totalStudents*100)
	fmt.Printf("Fail: %.2f%%\n", failPercent)
	fmt.Printf("Average: %.2f", averageGrade)
}
