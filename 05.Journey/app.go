package main

import "fmt"

func main() {

	var budget, spendMoney float64
	var season, destination, place string

	fmt.Scan(&budget, &season)

	if budget <= 100 {
		destination = "Bulgaria"

		switch season {
		case "summer":
			spendMoney = budget * 0.3
			place = "Camp"
		case "winter":
			spendMoney = budget * 0.7
			place = "Hotel"
		}
	} else if budget <= 1000 {
		destination = "Balkans"

		switch season {
		case "summer":
			spendMoney = budget * 0.4
			place = "Camp"
		case "winter":
			spendMoney = budget * 0.8
			place = "Hotel"
		}
	} else if budget > 1000 {
		destination = "Europe"
		spendMoney = budget * 0.9
		place = "Hotel"
	}
	fmt.Printf("Somewhere in %s\n", destination) //слагаме \n, за да принтира на два реда
	fmt.Printf("%s - %.2f", place, spendMoney) 
}