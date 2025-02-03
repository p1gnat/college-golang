package main

import "fmt"

func main() {
	var num int

	fmt.Print("Введите число , чтобы начать отсчёт: ")
	fmt.Scan(&num)

	for i := 1; i < num + 1; i++ {
		fmt.Printf("%v\n", i)
	}
}