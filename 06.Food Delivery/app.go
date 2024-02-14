package main

import "fmt"

func main() {

	//const chickenMenue float32 = 10.35
	//const fishMenue float32 = 12.40
	//const vegieMenue float32 = 8.15
	//const deliveryPrice float32 = 2.50

	var chichkenMenueCount, fishMenueCount, vegieMenueCount int

	fmt.Scanln(&chichkenMenueCount)
	fmt.Scanln(&fishMenueCount)
	fmt.Scanln(&vegieMenueCount)

	var totalPrice float32 = (float32(chichkenMenueCount) * 10.35) + (float32(fishMenueCount) * 12.40) + (float32(vegieMenueCount) * 8.15)

	var desertPrice float32 = totalPrice * 0.20

	var finalPrice = totalPrice + desertPrice + 2.50

	fmt.Println(finalPrice)
}