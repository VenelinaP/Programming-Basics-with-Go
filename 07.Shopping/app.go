package main

import "fmt"

func main() {

	var peterBudget float64
	var videocardsCount, processorsCount, ramCount int

	fmt.Scanln(&peterBudget)
	fmt.Scanln(&videocardsCount)
	fmt.Scanln(&processorsCount)
	fmt.Scanln(&ramCount)

	allVideocardPrice := float64(videocardsCount) * 250
	allProcesorPrice := float64(processorsCount) * (allVideocardPrice * 0.35)
	allRamPrice := float64(ramCount) * (allVideocardPrice * 0.10)

	finalPrice := allVideocardPrice + allProcesorPrice + allRamPrice

	if videocardsCount > processorsCount {
		finalPrice *= 0.85
	}

	if finalPrice <= peterBudget {
		leftBudget := peterBudget - finalPrice
		fmt.Printf("You have %.2f leva left!", leftBudget)
	} else {
		neededMoney := finalPrice - peterBudget
		fmt.Printf("Not enough money! You need %.2f leva more!", neededMoney)
	}
}
