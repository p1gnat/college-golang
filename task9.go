package main

import (
	"fmt"
	"strings"
)

func main() {
	var maxNumber int = 14
	
	for i := 1; i < maxNumber; i += 2 {
		spaces := strings.Repeat(" ", (maxNumber - i) / 2)
		fmt.Print(spaces)

		for j := i; j > 0; j-- {
			fmt.Print("*")
		}
		
		fmt.Print("\n")
	}
}