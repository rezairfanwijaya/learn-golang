package main

import "fmt"

func factorial(n int) int {
	if n > 0 {
		return n * factorial(n-1)
	}
	return 1
}

func main() {
	a := factorial(3)
	fmt.Println(a)
}
