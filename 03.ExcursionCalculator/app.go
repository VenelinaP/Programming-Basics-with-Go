package main

import "fmt"

func main() {

	var people int
	var season string
	var pricePerOne float64

	fmt.Scanln(&people)
	fmt.Scanln(&season)

	switch season {
	case "spring":
		if people <= 5 {
			pricePerOne = 50.00
		} else if people > 5 {
			pricePerOne = 48.00
		}
	case "summer":
		if people <= 5 {
			pricePerOne = 48.50
		} else if people > 5 {
			pricePerOne = 45.00
		}
	case "autumn":
		if people <= 5 {
			pricePerOne = 60.00
		} else if people > 5 {
			pricePerOne = 49.50
		}
	case "winter":
		if people <= 5 {
			pricePerOne = 86.00
		} else if people > 5 {
			pricePerOne = 85.00
		}
	}
	if season == "summer" {
		pricePerOne *= 0.85
	} else if season == "winter" {
		pricePerOne *= 1.08
	}
	 
	var finalPrice float64 = pricePerOne * float64(people)
	fmt.Printf("%.2f leva.", finalPrice)
}

