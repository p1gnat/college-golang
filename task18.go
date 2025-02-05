package main

import "fmt"

func main () {
	var str string

	fmt.Print("Введите строку для переворота: ")
	fmt.Scanln(&str)
	fmt.Print("Перевёрнутая строка: ")

	for i := len([]rune(str)) - 1; i >= 0; i-- {
		fmt.Print(string([]rune(str)[i]))
	}

	fmt.Println()
}
