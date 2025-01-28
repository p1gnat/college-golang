package main

import "fmt"

func main () {
	var num int

	fmt.Print("Введите ваш число: ")
	fmt.Scan(&num)
	if num > 0 {
		fmt.Printf("Число %d - положительное\n", num)
	} else if num < 0 {
		fmt.Printf("Число %d - отрицательное\n", num)
	} else {
		fmt.Printf("Число %d - является нулём\n", num)
	}
}