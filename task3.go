package main

import "fmt"

func main() {
	var num int

	fmt.Print("Введите число для таблицы умножения: ")
	fmt.Scan(&num)

	for i := 1; i <= 10; i++ {
		fmt.Printf("%d * %d = %d\n", num, i, num * i)
	}
}