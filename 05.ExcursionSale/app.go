package main

import "fmt"

func main() {

	var seaCount, mountainCount, profit int
	var input string

	fmt.Scanln(&seaCount)
	fmt.Scanln(&mountainCount)

	for {
		fmt.Scanln(&input)

		if input == "Stop" || (seaCount == 0 && mountainCount == 0) {
			break
		}

		if input == "sea" && seaCount > 0 {
			profit += 680
			seaCount--
		} else if input == "mountain" && mountainCount > 0 {
			profit += 499
			mountainCount--
		}
	}
	if seaCount == 0 && mountainCount == 0 {
		fmt.Println("Good job! Everything is sold.")
		fmt.Printf("Profit: %d leva.", profit)

	} else {
		fmt.Printf("Profit: %d leva.\n", profit)
	}
}