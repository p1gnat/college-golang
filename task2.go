package main

import "fmt"

func main() {
	var num1 int
	var num2 int

	fmt.Print("Введите первое число: ")
	fmt.Scan(&num1)
	fmt.Print("Введите второе число: ")
	fmt.Scan(&num2)

	sum := num1 + num2
	fmt.Printf("Сумма чисел: %d\n", sum)
}