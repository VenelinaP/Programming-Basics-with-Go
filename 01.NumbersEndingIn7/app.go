package main

import "fmt"

func main() {

	for n := 1; n <= 1000; n++ {

		if n%10 == 7 {
			fmt.Println(n)
		}
	}
}