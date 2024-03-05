package main

import (
	"fmt"
	"math"
)

func main() {

	var seriesTitle string
	var timeOfSeries, lunchBreak int

	fmt.Scanln(&seriesTitle)
	fmt.Scanln(&timeOfSeries)
	fmt.Scanln(&lunchBreak)

	timeForLunch := float64(lunchBreak) / 8.0 //за да изчисли времето за почивка във float64
	timeForRest := float64(lunchBreak) / 4.0

	allTime := float64(timeOfSeries) + timeForLunch + timeForRest
	
	if allTime <= float64(lunchBreak){
		timeLeft := math.Ceil(float64(lunchBreak) - allTime)
		fmt.Printf("You have enough time to watch %s and left with %.0f minutes free time.", seriesTitle, timeLeft) //%.0f - за да няма знаци след десетичната запетая
	} else {
		neededTime := math.Ceil(allTime - float64(lunchBreak))
		fmt.Printf("You don't have enough time to watch %s, you need %.0f more minutes.", seriesTitle, neededTime)
	}
}