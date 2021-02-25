package main

import (
	"fmt"
)

func m() {
	fmt.Println(factorial(3))
	fmt.Print(8 % 10)
}

func factorial(x int) int {
	if x == 1 {
		return 1
	}
	return x * factorial(x-1)
}
