package main

import "fmt"

func isPalindrom (str string) bool {
	var check string = ""
	for i := len(str) - 1; i >= 0; i-- {
		check += string(str[i])
	} 

	return check == str
}

func main() {
	fmt.Printf("isPalindrom('abba') = %v\n", isPalindrom("abba"))
	fmt.Printf("isPalindrom('abbas') = %v\n", isPalindrom("abbas"))
	fmt.Printf("isPalindrom('101') = %v\n", isPalindrom("101"))
}