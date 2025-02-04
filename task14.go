package main

import "fmt"

// Написать функцию, которая принимает два числа и возвращает их произведение.

func sum (num1 int, num2 int) int {
	return num1 + num2
}

func main() {
	fmt.Printf("sum(200, 92) = %v\n", sum(200, 92))
	fmt.Printf("sum(980890, 12) = %v\n", sum(980890, 12))
	fmt.Printf("sum(100, 0) = %v\n", sum(100, 0))
	fmt.Printf("sum(210, 78987) = %v\n", sum(210, 78987))
}