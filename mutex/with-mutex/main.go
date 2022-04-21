package main

import (
	"fmt"
	"sync"
)

// untuk mengatasi masalah seperti pada folder no-mutex maka kita harus melakukan beberapa settingan seperti pada stuct dan function add

type Num struct {
	val int
	sync.Mutex
}

func (n *Num) add() {
	// tambah lock agar tidak bisa diakses oleh goroutine lain
	n.Lock()
	// increment
	n.val++
	// hapus lock agar bisa diakses oleh goroutine lain
	n.Unlock()
}

func (n *Num) show() (res int) {
	res = n.val
	return
}

func main() {
	var wg sync.WaitGroup
	var mtx sync.Mutex
	var num Num

	for i := 0; i < 3000; i++ {
		// add wg
		wg.Add(1)
		for i := 0; i < 3000; i++ {
			go func() {
				// lock agar tidak bisa diakses oleh goroutine lain
				mtx.Lock()
				// increment
				num.add()
				// hapus lock agar bisa diakses oleh goroutine lain
				mtx.Unlock()
			}()
		}
		wg.Done()
	}

	wg.Wait()
	fmt.Println("nilai:", num.show())
	// oututnya sudah mendekarti 9000000
}
