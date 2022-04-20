package main

import (
	"fmt"
	"runtime"
)

// buffer yang dimaksud adalah batasan daya yang bisa diisi oleh sebuah channel
func main() {

	runtime.GOMAXPROCS(1)

	// buat channel dengan buffer 2 yang artinya channel ini dapat menampung data sebanyak 3 (0,1,2)
	myChannel := make(chan int, 2)

	// go routine untuk menerima data dari channel
	go func() {
		for {
			// masukkan data dari channel ke variabel data i
			i := <-myChannel
			fmt.Println("Data yang diterima:", i)
		}
	}()

	// loop untuk mengisi data ke channel
	// meskipun perulangan dilakukan sebnyak 5 kali namun tetap data yang dapat ditampung adalah 3 (buffer + 1)
	for i := 0; i < 5; i++ {
		fmt.Println("Data yang dikirim:", i)
		myChannel <- i
	}

}
