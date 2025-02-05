package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main () {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите строку для замены пробелов: ")

	str, _ := reader.ReadString('\n')
	str = strings.TrimSpace(str)

	result := strings.ReplaceAll(string([]rune(str)), " ", ".")
	fmt.Printf("Заменённая строка без пробелов: %v\n", result)
}