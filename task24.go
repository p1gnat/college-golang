package main

import "fmt"

func IntAbs(value int) int {
	if value < 0 {
		return -value
	}
	return value
}

func GCD(a, b int) int {
	for b != 0 {
		a %= b
		if a == 0 {
			return IntAbs(b)
		}
		b %= a
	}
	return IntAbs(a)
}

func main() {
	fmt.Println("GCD(10,100): ", GCD(10,100))
	fmt.Println("GCD(1,76): ", GCD(12,76))
	fmt.Println("GCD(5, 83): ", GCD(5, 80))
	fmt.Println("GCD(237549, 4324): ", GCD(237549, 4324))
}