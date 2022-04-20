package main

import "fmt"

// channel range close ini akan berfungsi menonaktifkan channel ketika masuk di perulangan for

func sendMessage(ch chan string) {
	for i := 0; i < 20; i++ {
		ch <- fmt.Sprintf("data ke: %d\n", i)
	}
	// tutup channel setelah perulangan selesai
	close(ch) // kalau ini di comment maka akan error deadlock
}

func printMessage(ch chan string) {
	for message := range ch {
		fmt.Println(message)
	}
}

func main() {
	var msg = make(chan string)
	go sendMessage(msg)
	printMessage(msg)
}
