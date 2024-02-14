package main

import "fmt"

func main() {

	var l, w, h int
	var percent float32

	fmt.Scanln(&l)
	fmt.Scanln(&w)
	fmt.Scanln(&h)
	fmt.Scanln(&percent) 

	var tank int = l * w * h
	var tankVolume float32 = float32(tank) * 0.001 //tankVolume float32 = float32(tank) / 1000
		
	var neededSpace float32 = tankVolume * (1 - (percent / 100))
	
	fmt.Println(neededSpace)
}