package main

import (
	"fmt"
	"math"
)

func main() {

	var recordInSec, distanceInMetres, secondsForOneMeter float64

	fmt.Scanln(&recordInSec)
	fmt.Scanln(&distanceInMetres)
	fmt.Scanln(&secondsForOneMeter)

	timeForSwiming := distanceInMetres * secondsForOneMeter
	delay := math.Floor(distanceInMetres /15) * 12.5

	finalTime := timeForSwiming + delay
	
	if finalTime < recordInSec {
		fmt.Printf("Yes, he succeeded! The new world record is %.2f seconds.", finalTime)
	} else {
		neededTime := finalTime - recordInSec 
		fmt.Printf("No, he failed! He was %.2f seconds slower.",neededTime)
	}
}