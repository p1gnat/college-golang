package main

import "fmt"

func main() {
	for i := 1; i < 101; i++ {
		if i % 2 == 0 {
			fmt.Printf("%v\n", i)
		}
	}
}