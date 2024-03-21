package main

import "fmt"

func main() {

	var tabs, salary int
	var tabOpen string

	fmt.Scanln(&tabs)
	fmt.Scanln(&salary)

	for i := 0; i < tabs; i++ {
		fmt.Scanln(&tabOpen)

		switch tabOpen {
		case "Facebook":
			salary -= 150

		case "Instagram":
			salary -= 100

		case "Reddit":
			salary -= 50
		}

		if salary <= 0 { //така прекъсваме цикъла, защото парите са 0
			break        // трябва да е в иф проверка винаги
		}
	}

	if salary > 0 {
		fmt.Println(salary)
	} else {
		fmt.Println("You have lost your salary.")
	}
}
