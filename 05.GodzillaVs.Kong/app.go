package main

import (
	"fmt"
	"math"
)

func main() {
	var budget float64
	var statists int
	var clothesPrice float64

	fmt.Scanln(&budget)
	fmt.Scanln(&statists)
	fmt.Scanln(&clothesPrice)

	decor := budget * 0.10

	if statists > 150 {
		clothesPrice *= 0.90
	}

	clothes := float64(statists) * clothesPrice
	allMoney := decor + clothes
	moneyNeeded := math.Abs(budget - allMoney)

	if allMoney > budget {
		fmt.Println("Not enough money!")
		fmt.Printf("Wingard needs %.2f leva more.\n", moneyNeeded)
	} else {
		fmt.Println("Action!")
		fmt.Printf("Wingard starts filming with %.2f leva left.\n", moneyNeeded)
	}
}