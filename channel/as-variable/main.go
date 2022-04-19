package main

import (
	"fmt"
)

func main() {
	// inisiasi channel menggunakan make (chan, tipe data)
	var pesan = make(chan string)

	// bikin closure function
	var sayHello = func(nama string) {
		data := fmt.Sprintf("Hello %s", nama)
		// masukan nilai ke channel
		// nanti pesan akan berbunyi "Hello + nama"
		pesan <- data
	}

	// bikin goroutine
	go sayHello("Budi")
	go sayHello("Andi")
	go sayHello("Siti")

	// keluarkan data dari channel yang telah diisi sebelumnnya di func sayHello

	pesan1 := <-pesan
	fmt.Println(pesan1)

	pesan2 := <-pesan
	fmt.Println(pesan2)

	pesan3 := <-pesan
	fmt.Println(pesan3)

}
