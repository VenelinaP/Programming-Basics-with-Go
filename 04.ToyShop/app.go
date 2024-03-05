package main

import (
	"fmt"
	"math"
)

func main() {

	const puzzelPrice float64 = 2.60
	const dollsPrice float64 = 3
	const bearsPrice float64 = 4.10
	const minionPrice float64 = 8.20
	const truckPrice float64 = 2

	var excursionPrice float64
	var puzzelsCount, dollsCount, bearsCount, minionsCount, truckCount int

	fmt.Scanln(&excursionPrice)
	fmt.Scanln(&puzzelsCount)
	fmt.Scanln(&dollsCount)
	fmt.Scanln(&bearsCount)
	fmt.Scanln(&minionsCount)
	fmt.Scanln(&truckCount)

	sumPrices := float64(puzzelsCount)*puzzelPrice + float64(dollsCount)*dollsPrice + float64(bearsCount)*bearsPrice + float64(minionsCount)*minionPrice + float64(truckCount)*truckPrice
	
	alltoys := puzzelsCount + dollsCount + bearsCount + minionsCount + truckCount

	if alltoys >= 50 {
		sumPrices *= 0.75
	}

	rent := sumPrices * 0.1

	finalIncome := sumPrices - rent
	diff := math.Abs(finalIncome - excursionPrice)

	if finalIncome >= excursionPrice {
		//moneyLeft := finalIncome - excursionPrice

		fmt.Printf("Yes! %.2f lv left.", diff) //moneyLeft )
	} else {
		//moneyNeeded := excursionPrice - finalIncome

		fmt.Printf("Not enough money! %.2f lv needed.", diff) //moneyNeeded)
	}
}
