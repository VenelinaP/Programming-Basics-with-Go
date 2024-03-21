package main

import "fmt"

func main() {

	var num int

	fmt.Scanln(&num)

	var isValid bool = num >= 100 && num <= 200 || num == 0

	if !isValid {
		fmt.Println("invalid")
	} else {
		fmt.Println()
	}
}