package main

import "fmt"

func main() {
	var userArray [5]int
	var sum int

	fmt.Println("Введите 5 чисел:")
	
	for i := 0; i < 5; i++ {
		fmt.Printf("%v Число: ", i + 1)
		fmt.Scan(&userArray[i])
		sum += userArray[i]
	}

	fmt.Printf("\n")

	fmt.Println("Массив:", userArray[:])
	fmt.Println("Сумма введённых чисел:", sum)
}