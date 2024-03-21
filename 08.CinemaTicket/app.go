package main

import "fmt"

func main() {

	var day string
	var ticketPrice int

	fmt.Scanln(&day)

	if day == "Monday" || day == "Tuesday" || day == "Friday" {
		ticketPrice = 12
	} else if day == "Wednesday" || day == "Thursday" {
		ticketPrice = 14
	} else if day == "Saturday" || day == "Sunday" {
		ticketPrice = 16
	}
	fmt.Println(ticketPrice)
}