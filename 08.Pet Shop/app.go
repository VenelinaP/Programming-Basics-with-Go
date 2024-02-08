package main

import "fmt"

func main() {

	var dogsFood int
	var catsFood int

	fmt.Scanln(&dogsFood)
	fmt.Scanln(&catsFood)

	var allDogsFood float32 = 2.50 * float32(dogsFood)
	var allCatsFood float32 = 4 * float32(catsFood)

	var finalPrice float32 = allDogsFood + allCatsFood

	fmt.Printf("%f lv.", finalPrice)
}