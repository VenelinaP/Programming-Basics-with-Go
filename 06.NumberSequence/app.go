package main

import (
	"fmt"
	"math"
)

func main() {

	var num, currNum int
	fmt.Scanln(&num)

	var minNum int32 = math.MaxInt32 //връща максималното инт
	var maxNum int32 = math.MinInt32 //връща минималното число

	for i := 0; i < num; i++ {
		fmt.Scanln(&currNum)

		if int32(currNum) < minNum {
			minNum = int32(currNum)
		}

		if int32(currNum) > maxNum {
			maxNum = int32(currNum)
		}
	}
	fmt.Printf("Max number: %d\n", maxNum)
	fmt.Printf("Min number: %d", minNum)
}
