package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func createFileAndSum () {
	fmt.Println("Запись файла...")
	file, _ := os.Create("text_file_task21.txt")
	defer file.Close()

	numbers := []string{"10", "20", "30", "40", "3"}
	for _, num := range numbers {
		file.WriteString(num + "\n")
	}

	file, _ = os.Open("text_file_task21.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		sum += num
	}

	fmt.Println("Сумма чисел в файле: ", sum)
}

func main () {
	createFileAndSum()
}