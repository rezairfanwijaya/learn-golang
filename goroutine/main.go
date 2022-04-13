package main

import (
	"fmt"
	"runtime"
)

func Hello(loop int, kata string) {
	for i := 0; i < loop; i++ {
		fmt.Println(i, kata)
	}
}

func main() {

	runtime.GOMAXPROCS(1)
	fmt.Println("dengan goroutine")
	go Hello(5, "goroutine")
	// Hello(5, "im")
	// Hello(5, "goroutine")

	// ini tidak menggunakan goroutine
	fmt.Println("no goroutine")
	Hello(5, "Hello World")

	var input string
	fmt.Scanln(&input)

	fmt.Println("input:", input)
}
