package main

import "fmt"

func main() {

	var age, toyPrice, toyCounter, money int
	var washerPrice float64

	fmt.Scanln(&age)
	fmt.Scanln(&washerPrice)
	fmt.Scanln(&toyPrice)

	for i := 1; i <= age; i++ {

		if i % 2 != 0 {
			toyCounter++		
		} else {
			money += i * 5
			money--
		}		
	}
	money += toyCounter * toyPrice

	if float64(money) >= (washerPrice) {
		moneyLeft := float64(money) - washerPrice

		fmt.Printf("Yes! %.2f", moneyLeft)
	} else {
		moneyNeeded := washerPrice - float64(money)

		fmt.Printf("No! %.2f", moneyNeeded)
	}
}