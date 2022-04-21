package main

import (
	"fmt"
	"sync"
)

// waitgroup diguakan untuk mensinkronisasi antar goroutine

// Print digunakan untuk mencetak data
// wg digunakan untuk menghentikan proses
func Print(wg *sync.WaitGroup, msg string) {
	// wg.Done ini akan mengurangi jumlah goroutine yang sedang menunggu
	defer wg.Done()
	fmt.Println(msg)
}

func main() {
	// inisiate waitgroup
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		// wg.Add digunakan untuk menambahkan jumlah goroutine yang sedang menunggu (jumlah goroute yang akan dijalankan)
		wg.Add(1)
		var data = fmt.Sprintln("Data ke", i)
		go Print(&wg, data)
	}

	// wg.Wait digunakan untuk menunggu sampai semua goroutine selesai
	wg.Wait()
}
