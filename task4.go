package main

import "fmt"

func main () {
	var num int

	fmt.Print("Введите число для проверки чётности: ")
	fmt.Scan(&num)
	if num % 2 == 0 {
		fmt.Printf("%d Чётное\n", num)
	} else {
		fmt.Printf("%d Не чётное\n", num)
	}
}