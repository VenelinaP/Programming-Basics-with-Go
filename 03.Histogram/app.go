package main

import "fmt"

func main() {

	var n, currNum, p1, p2, p3, p4, p5 int
	fmt.Scanln(&n)

	for i := 0; i < n; i++ {
		fmt.Scanln(&currNum)

		if currNum < 200 {
			p1++
		} else if currNum <= 399 {
			p2++
		} else if currNum <= 599 {
			p3++
		} else if currNum <= 799 {
			p4++
		} else if currNum >= 800 {
			p5++
		}
	}

	percentP1 := float32(p1) / float32(n) * 100
	percentP2 := float32(p2) / float32(n) * 100
	percentP3 := float32(p3) / float32(n) * 100
	percentP4 := float32(p4) / float32(n) * 100
	percentP5 := float32(p5) / float32(n) * 100

	fmt.Printf("%.2f%%\n", percentP1)
	fmt.Printf("%.2f%%\n", percentP2)
	fmt.Printf("%.2f%%\n", percentP3)
	fmt.Printf("%.2f%%\n", percentP4)
	fmt.Printf("%.2f%%", percentP5)
}