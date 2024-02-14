package main

import "fmt"

func main() {

	const nylonPrice float32 = 1.50
	const paintPrice float32 = 14.50
	const tinerPrice float32 = 5.00
	const bags float32 = 0.40

	var nylon, paint, tiner, hours int

	fmt.Scanln(&nylon)
	fmt.Scanln(&paint)
	fmt.Scanln(&tiner)
	fmt.Scanln(&hours)

	var allNylon = (float32(nylon) + 2) * nylonPrice
	var allPaint = float32(paint) * 1.1 * paintPrice
	var allTiner = float32(tiner) * tinerPrice

	var totalPrice = allNylon + allPaint + allTiner + bags

	var pricePerHour float32 = totalPrice * 0.30
	var priceForWorkers float32 = float32(hours) * pricePerHour

	var finalPrice = totalPrice + priceForWorkers

	fmt.Println(finalPrice)
}