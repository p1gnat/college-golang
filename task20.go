package main

import (
	"fmt"
	"os"
)

func createFile () {
	fmt.Println("Запись файла...")
	file, _ := os.Create("text_file.txt")
	defer file.Close()

	file.WriteString("Задание 20")

	content, _ := os.ReadFile("text_file.txt")

	fmt.Println("Содержимое файла: ", string(content))
}

func main () {
	createFile()
}