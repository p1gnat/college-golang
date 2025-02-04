package main

import "fmt"

func findFactorial (num int) int {
	var factorial int = 1
	for i := 1; i <= num; i++ {
		factorial *= i
	}
	return factorial
}

func main() {
	fmt.Printf("findFactorial(19) = %v\n", findFactorial(19))
	fmt.Printf("findFactorial(9) = %v\n", findFactorial(9))
	fmt.Printf("findFactorial(5) = %v\n", findFactorial(5))
}