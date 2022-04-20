package main

import (
	"fmt"
	"math/rand"
	"time"
)

// timeout channel ini bisa mengatur kapan channel akan menerima data

func sendData(ch chan<- string) {
	for i := 0; true; i++ {
		ch <- fmt.Sprintf("Data ke %d", i)
		// set format time untuk mengatur timeout
		time.Sleep(time.Duration(rand.Int()%10+1) * time.Second)
	}
}

func reveiveData(ch <-chan string) {

	for {
		select {
		case data := <-ch:
			fmt.Println(data)
		case <-time.After(time.Second * 3):
			fmt.Println("Timeout : tidak ada aktivitas dalam 3 detik")
			break // break untuk menghentikan looping
		}
	}
}

func main() {
	var ch = make(chan string)
	go sendData(ch)
	reveiveData(ch)
}
