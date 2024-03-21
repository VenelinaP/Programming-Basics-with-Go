package main

import "fmt"

func main() {

	var num1, num2, result int
	var operator string

	fmt.Scan(&num1, &num2, &operator)

	switch operator {
	case "+", "-", "*":
		if operator == "+" {
			result = num1 + num2
		} else if operator == "-" {
			result = num1 - num2
		} else if operator == "*" {
			result = num1 * num2
		}
		
		if result % 2 == 0 {
			fmt.Printf("%d %s %d = %d - even", num1, operator, num2, result)
		} else {
			fmt.Printf("%d %s %d = %d - odd", num1, operator, num2, result)
		}
	case "/":
		if num2 == 0 {
			fmt.Printf("Cannot divide %d by zero", num1)
		} else {
			resultAfterDivide := float64(num1) / float64(num2)
			fmt.Printf("%d / %d = %.2f", num1, num2, resultAfterDivide)
		}		
	case "%":
		if num2 == 0 {
			fmt.Printf("Cannot divide %d by zero", num1)
		} else {
			result = num1 % num2
			fmt.Printf("%d %% %d = %d", num1, num2, result) //слагаме %% за да се изпише правилно в изхода
		}		
	}
}