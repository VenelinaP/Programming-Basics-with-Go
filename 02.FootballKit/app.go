package main

import "fmt"

func main() {

	var shirt float64
	var sum float64

	fmt.Scanln(&shirt)
	fmt.Scanln(&sum)

	var shorts float64 = shirt * 0.75
	var socks float64 = shorts * 0.20
	var shoes float64 = 2 * (shirt + shorts)

	var totalPrice float64 = shirt + shorts + socks + shoes
	var discountPrice float64 = totalPrice * 0.85

	if discountPrice >= sum {
		fmt.Printf("Yes, he will earn the world-cup replica ball!\n")
		fmt.Printf("His sum is %.2f lv.", discountPrice)
	} else {
		var moneyNeeded float64 = sum - discountPrice
		fmt.Printf("No, he will not earn the world-cup replica ball.\n")
		fmt.Printf("He needs %.2f lv. more.", moneyNeeded)
	}
}