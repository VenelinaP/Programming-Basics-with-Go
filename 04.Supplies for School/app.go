package main

import "fmt"

func main() {

	const pensPrice float32 = 5.80 //променлива за цената, това е константно условие, затова го създаваме отначало
	const marckersPtice float32 = 7.20
	const cleanerPrice float32 = 1.20

	var pens int //всички променливи от един тип могат да се изпишат на един ред
	var marckers int
	var cleaner int
	var percentDiscount int

	fmt.Scanln(&pens)
	fmt.Scanln(&marckers)
	fmt.Scanln(&cleaner)
	fmt.Scanln(&percentDiscount)

	var allPens float32 = float32(pens) * pensPrice
	var allMarckers float32 = float32(marckers) * marckersPtice
	var allCleaner float32 = float32(cleaner) * cleanerPrice
	
	var totalPrice = allPens + allMarckers + allCleaner
	var discount = totalPrice * float32(percentDiscount) / 100
	var finalPrice = totalPrice - discount

	fmt.Println(finalPrice)
}