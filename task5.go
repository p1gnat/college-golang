package main

import "fmt"

func main () {
	var num int

	fmt.Print("Введите ваш возраст: ")
	fmt.Scan(&num)
	if num >= 18 {
		fmt.Printf("Ваш возраст %d - Вы являетесь взрослым\n", num)
	} else {
		fmt.Printf("Ваш возраст %d - Вы не являетесь взрослым\n", num)
	}
}