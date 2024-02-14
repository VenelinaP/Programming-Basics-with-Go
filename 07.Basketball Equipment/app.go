package main

import "fmt"

func main() {

	var anualFee int

	fmt.Scanln(&anualFee)

	var shoes float32 = float32(anualFee) * 0.60
	var clothes float32 = shoes * 0.80
	var ball float32 = clothes / 4 //защото е 1/4 от цената а топката
	var accessories float32 = ball / 5

	var totalPrice float32 = float32(anualFee) + shoes + clothes + ball + accessories

	fmt.Println(totalPrice)

}