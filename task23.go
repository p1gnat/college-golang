package main

import "fmt"

func main () {
	arr := []int{100, 10000, 54, 92, 1000}
	fmt.Println("Массив для перебора: ", arr)
	minimum := arr[0]
	for i := 0; i <= len(arr) - 1; i++ {
		if arr[i] < minimum {
			minimum = arr[i]
		}
	}
	fmt.Println("Минимальное число из массива: ", minimum)
}







// package main

// import "fmt"


// func main() {
// 	arr := []int{10, 2, 3, 4, 5, 6, 7}
// 	minimum := arr[0]
// 	for i := 0; i <= len(arr) - 1; i++ {
// 		if arr[i] < minimum {
// 			minimum = arr[i]
// 		}
// 	}
// 	fmt.Println(minimum)
// }