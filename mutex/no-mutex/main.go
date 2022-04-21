package main

import (
	"fmt"
	"sync"
)

// disini kita akan coba membuat program yang bermasalah jika tidak memiliki mutex
// kalau begini berarti ada kesalahan pada program yang dinamakan race condition
// ini bisa dicek dengan cara go run -race main.go

// program ini akan mengincrement nilai sebanyak 3000 x 3000 kali

type Num struct {
	val int
}

// func untuk increment
func (n *Num) add() {
	n.val++
}

// func untuk mengambil nilai val
func (n *Num) show() (res int) {
	res = n.val
	return
}

func main() {
	var wg sync.WaitGroup
	var num Num

	for i := 0; i < 3000; i++ {
		// daftarkan semua goroutine
		wg.Add(1)

		// goroutine untuk increment
		go func() {
			for j := 0; j < 3000; j++ {
				num.add()
			}
			// menghentikan goroutine yang mengincrement
			wg.Done()
		}()
	}

	// menunggu semua goroutine selesai
	wg.Wait()

	// show data
	fmt.Println("nilai:", num.show())
	// harus nya data yg dihasilkan adalah 9000000
}
