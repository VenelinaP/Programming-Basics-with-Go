package main

import "fmt"

func main() {

	var n, currNum, sum int
	fmt.Scanln(&n)

	for i := 1; i <= n; i++ {
		fmt.Scanln(&currNum) //за да чете ново число от конзолата 

		sum += currNum				
	}

	fmt.Println(sum)
}