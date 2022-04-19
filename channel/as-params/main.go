package main

import "fmt"

func main() {
	var pesan = make(chan string)
	var nama = []string{"Budi", "Andi", "Siti"}

	for _, each := range nama {
		go func(nama string) {
			data := fmt.Sprintf("Hello %s", nama)
			pesan <- data
		}(each)
	}

	for i := 0; i < 3; i++ {
		fmt.Println(<-pesan)
	}
}
