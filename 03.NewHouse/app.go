package main

import "fmt"

func main() {

	const rosesPrice float64 = 5
	const dahliasPrice float64 = 3.80
	const tulipsPrice float64 = 2.80
	const narcissusPrice float64 = 3
	const gladiolusPrice float64 = 2.50

	var flowerType string
	var quantity, budget int
	var totalPrice float64

	fmt.Scan(&flowerType, &quantity, &budget)

	if flowerType == "Roses" {
		totalPrice = float64(quantity) * rosesPrice

		if quantity > 80 {
			totalPrice *= 0.9
		}

	} else if flowerType == "Dahlias" {
		totalPrice = float64(quantity) * dahliasPrice

		if quantity > 90 {
			totalPrice *= 0.85
		}

	} else if flowerType == "Tulips" {
		totalPrice = float64(quantity) * tulipsPrice

		if quantity > 80 {
			totalPrice *= 0.85
		}

	} else if flowerType == "Narcissus" {
		totalPrice = float64(quantity) * narcissusPrice

		if quantity < 120 {
			totalPrice *= 1.15
		}

	} else if flowerType == "Gladiolus" {
		totalPrice = float64(quantity) * gladiolusPrice

		if quantity < 80 {
			totalPrice *= 1.2
		}
	}
	if float64(budget) >= totalPrice {
		moneyLeft := float64(budget) - totalPrice
		fmt.Printf("Hey, you have a great garden with %d %s and %.2f leva left.", quantity, flowerType, moneyLeft)
	} else {
		moneyNeeded := totalPrice - float64(budget)
		fmt.Printf("Not enough money, you need %.2f leva more.", moneyNeeded)
	}
}
