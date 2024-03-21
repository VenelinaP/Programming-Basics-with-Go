package main

import "fmt"

func main() {

	var tournamentCount, startingPoints, tournamentPoints, winTournaments int
	var result string

	fmt.Scanln(&tournamentCount)
	fmt.Scanln(&startingPoints)

	for i := 0; i < tournamentCount; i++ {
		fmt.Scanln(&result)

		switch result {
		case "W":
			tournamentPoints += 2000
			winTournaments++
		case "F":
			tournamentPoints += 1200		
		case "SF":
			tournamentPoints += 720
		}		
	}

	averagePoints := tournamentPoints / tournamentCount 

	var percentWinTournament float64 = (float64(winTournaments)) / float64(tournamentCount) * 100
	var finalPoints = startingPoints + tournamentPoints
	
	fmt.Printf("Final points: %d\n", finalPoints)
	fmt.Printf("Average points: %d\n", averagePoints)
	fmt.Printf("%.2f%%", percentWinTournament)
}