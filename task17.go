package main

import (
	"bufio"
	"fmt"
	"os"
)

func contains (s string, r rune) bool {
	for _, v := range s {
		if v == r {
			return true
		}
	}
	return false
}

func main () {
	vowels := "аеёиоуыэюяАЕЁИОУЫЭЮЯ"
	var counter int

	fmt.Print("Введите строку с гласными: ")
	reader := bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')

	for _, char := range str {
		if contains(vowels, char) {
			counter++
		}
	}

	fmt.Printf("В строке было %v гласных\n", counter)
}
