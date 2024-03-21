package main

import (
	"fmt"
	"math"
)

func main() {

	var n, oddSum, evenSum, currNum int

	fmt.Scanln(&n)
	
	for i := 1; i <= n; i++ {
		fmt.Scanln(&currNum)

		if i % 2 == 1 {
			oddSum += currNum			
		} else {
			evenSum += currNum
		}		
	}

	if evenSum == oddSum {
		fmt.Println("Yes")
		fmt.Printf("Sum = %d", oddSum)		
	} else {
		diff := math.Abs(float64(evenSum) - float64(oddSum))

		fmt.Println("No")
		fmt.Printf("Diff = %.0f", diff)
	}
}