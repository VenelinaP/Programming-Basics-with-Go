package main

import "fmt"

func main() {

	var budget, fishers int
	var season string

	fmt.Scanln(&budget)
	fmt.Scanln(&season)
	fmt.Scanln(&fishers)

	var rent float64

	switch season {
	case "Spring":
		rent = 3000
	case "Summer", "Autumn":
		rent = 4200
	case "Winter":
		rent = 2600
	}
	if fishers <= 6 {
		rent *= 0.90	
	} else if fishers <= 11 {
		rent *= 0.85		
	} else {
		rent *= 0.75
	}
	
	if fishers % 2 == 0 && season != "Autumn" {
		rent *= 0.95
	}

	if float64(budget) >= rent {
		moneyLeft := float64(budget) - rent
		fmt.Printf("Yes! You have %.2f leva left.", moneyLeft)
	} else {
		moneyNeeded := rent - float64(budget)
		fmt.Printf("Not enough money! You need %.2f leva.", moneyNeeded)
	}	
}
	


