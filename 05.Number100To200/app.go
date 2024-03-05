package main

import "fmt"

func main() {

	var num int
	fmt.Scanln(&num)

	if num < 100 {
		fmt.Println("Less than 100")
	} else if num <= 200 {
		fmt.Println("Between 100 and 200")
	} else if num > 200 {
		fmt.Println("Greater than 200")
	}
}
