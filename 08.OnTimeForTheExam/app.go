package main

import "fmt"

func main() {

	var examHour, arrivalHour, examMinutes, arrivalMinutes int

	fmt.Scan(&examHour, &examMinutes, &arrivalHour, &arrivalMinutes)

	examTimeInMinutes := examHour*60 + examMinutes
	arrivalTimeInMinutes := arrivalHour*60 + arrivalMinutes

	if arrivalTimeInMinutes > examTimeInMinutes {
		fmt.Println("Late")

		diff := arrivalTimeInMinutes - examTimeInMinutes

		if diff < 60 {
			fmt.Printf("%d minutes after the start", diff)

		} else {
			lateHour := diff / 60
			lateMinutes := diff % 60

			fmt.Printf("%d:%.2d hours after the start", lateHour, lateMinutes) //%.2d, за да изпише 0 преди числото ако е едноцифрено
		}
	} else if examTimeInMinutes - arrivalTimeInMinutes <= 30 || examTimeInMinutes == arrivalTimeInMinutes {
		fmt.Println("On time")

		if examTimeInMinutes != arrivalTimeInMinutes {
			diff := examTimeInMinutes - arrivalTimeInMinutes

			fmt.Printf("%d minutes before the start", diff)
		}

	} else if examTimeInMinutes - arrivalTimeInMinutes > 30 {
		fmt.Println("Early")

		diff := examTimeInMinutes - arrivalTimeInMinutes

		if diff < 60 {
			fmt.Printf("%d minutes before the start", diff)
		} else {
			earlyHour := diff / 60
			earlyMinutes := diff % 60

			fmt.Printf("%d:%.2d hours before the start", earlyHour, earlyMinutes)
		}
	}
}
