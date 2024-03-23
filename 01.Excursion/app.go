package main

import "fmt"

func main() {

	var people, nights, cards, tickets int

	fmt.Scanln(&people)
	fmt.Scanln(&nights)
	fmt.Scanln(&cards)
	fmt.Scanln(&tickets)

	var nightPrice float64 = 20.00 * float64(nights)
	var cardPrice float64 = 1.60 * float64(cards)
	var ticketPrice float64 = 6.00 * float64(tickets)
	
	var priceForOne float64 = nightPrice + cardPrice + ticketPrice

	var priceForGroup float64 = float64(people) * priceForOne
	var additionalExpences float64 = (priceForGroup * 2.5) / 2

	fmt.Printf("%.2f", additionalExpences)
}