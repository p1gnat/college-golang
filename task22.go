package main

import "fmt"

func bubbleSort (arr []int) {
	length := len(arr)
	for i := 0; i < length - 1; i++ {
		for j := 0; j < length - i - 1; j++ {
			if arr[j] > arr[j + 1] {
				arr[j], arr[j + 1] = arr[j + 1], arr[j]
			}
		}
	}
}

func main() {
	arr := []int{5, 2, 9, 1, 5, 6}
	fmt.Println("Исходный массив:", arr)
	bubbleSort(arr)
	fmt.Println("Отсортированный массив:", arr)
}