package main

import (
	"fmt"
)

// direction channel ini memberikan hak akses ke channel
// dengan kata laim, kita bisa menentukan suatu channel berperan sebagai penerima/pengirim/both
/*
syntax:
ch chan int  == ini channel sebagai penerima sekaligus pengirim
ch chan<- int == ini channel sebagai pengirim saja
ch <- chan int == ini channel sebagai penerima saja
*/

func listSeller(ch chan<- string) {
	for i := 0; i < 30; i++ {
		ch <- fmt.Sprintf("Seller ke: %d\n", i+1)
	}
	close(ch)
}

func showSeller(ch <-chan string) {
	for seller := range ch {
		fmt.Println(seller)
	}
}

func main() {
	var seller = make(chan string)
	go listSeller(seller)
	showSeller(seller)
}
