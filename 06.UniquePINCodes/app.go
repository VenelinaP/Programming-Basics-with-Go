package main

import "fmt"

func main() {

	var firstNum, secondNum, thirdNum int

	fmt.Scanln(&firstNum)
	fmt.Scanln(&secondNum)
	fmt.Scanln(&thirdNum)

	for n1 := 2; n1 <= firstNum; n1 += 2 {
		for n2 := 2; n2 <= secondNum; n2++ {
			isPrime := true
			for i := 2; i*i <= n2; i++ {
				if n2 % i == 0 {
					isPrime = false
					break
				}
			}
			if isPrime {
				for n3 := 2; n3 <= thirdNum; n3 += 2 {
					fmt.Printf("%d %d %d\n", n1, n2, n3)
				}
			}
		}
	}
}
//Използването на i*i вместо i се използва за оптимизация.Когато проверяваме дали число n2 е просто, не е необходимо да проверяваме делителите му до самото число n2. 
//Вместо това, можем да проверим само до квадратния корен на n2, тъй като ако има делител по-голям от квадратния корен, то ще има и делител по-малък от квадратния корен.
//Така че, когато използваме i*i <= n2, ние избягваме излишни проверки и това оптимизира процеса на проверка за простота.
