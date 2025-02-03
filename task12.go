package main

import (
	"fmt"
	"slices"
)

func main() {
	var userArray [5]int
	var sum int

	fmt.Println("Введите 5 чисел для поиска наименьшего и наибольшего:")
	
	for i := 0; i < 5; i++ {
		fmt.Printf("%v Число: ", i + 1)
		fmt.Scan(&userArray[i])
		sum += userArray[i]
	}

	fmt.Printf("\n")

	userSlice := userArray[:]

	fmt.Println("Массив:", userArray[:])
	fmt.Println("Минимальное:", slices.Min(userSlice))
	fmt.Println("Максимальное:", slices.Max(userSlice))
}