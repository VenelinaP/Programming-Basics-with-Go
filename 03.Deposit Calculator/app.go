package main

import "fmt"

func main() {

	var depositSum float32
	var depositPeriod float32
	var anualInterest float32

	fmt.Scanln(&depositSum)
	fmt.Scanln(&depositPeriod)
	fmt.Scanln(&anualInterest)

	var anualSum float32 = (depositSum * (anualInterest / 100)) / 12
	var total float32 = depositSum + (depositPeriod * anualSum)

	fmt.Println(total)
}