package main

import "fmt"

func main() {

	var degrees int
	var timeOfTheDay, outfit, shoes string

	fmt.Scan(&degrees, &timeOfTheDay)

		if degrees >=10 && degrees <= 18 {
			switch timeOfTheDay {
			case "Morning":
			outfit = "Sweatshirt"
			shoes = "Sneakers"
			case "Afternoon", "Evening":
				outfit = "Shirt"
				shoes = "Moccasins"
			} 
		} else if degrees > 18 && degrees <= 24 {
			switch timeOfTheDay {
			case "Morning", "Evening":
				outfit = "Shirt"
				shoes = "Moccasins"
			case "Afternoon":
				outfit = "T-Shirt"
				shoes = "Sandals"
			} 
		} else if degrees > 24 {
			switch timeOfTheDay {
			case "Morning":
				outfit = "T-Shirt"
				shoes = "Sandals"
			case "Afternoon":
				outfit = "Swim Suit"
				shoes = "Barefoot"
			case "Evening":
				outfit = "Shirt"
				shoes = "Moccasins"
			}		
		}
	    fmt.Printf("It's %d degrees, get your %s and %s.", degrees, outfit, shoes)
}

