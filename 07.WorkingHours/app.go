package main

import "fmt"

func main() {

	var hour int
	var day, result string
	

	fmt.Scanln(&hour)
	fmt.Scanln(&day)

	if hour >= 10 && hour <= 18 {
		if day == "Monday" || day == "Tuesday" || day == "Wednesday" || day == "Thursday" || day == "Friday" || day == "Saturday" {
		result = "open"
		} else if day == "Sunday" {
			result = "closed"
		}
	} else {
		result = "closed"
		}
	fmt.Println(result)
}