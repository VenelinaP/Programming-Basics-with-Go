package main

import "fmt"

func main() {

	var animalType string
	fmt.Scanln(&animalType)

	switch animalType {
		case "dog":
			fmt.Println("mammal")
		case "crocodile", "tortoise", "snake":
			fmt.Println("reptile")
		default:
			fmt.Println("unknown")
	}
}