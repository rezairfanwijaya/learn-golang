package main

import "fmt"

// select ini berfungsi sama seperti switch case pada pengkondisian, jadi kita bisa memilih akan menggunakan channel mana pada kondisi tertentu.

// pada kasus kali ini kita akan menghitung rata-rata nilai dan juga nilai max dari integer of slice

// func untuk rata-rata nilai
func average(num []int, ch chan float64) {
	// kita bikin variable penampung hasil penjumlahan semua elemen dari slice
	sum := 0

	// looping element dan jumlahkan
	for _, val := range num {
		sum += val
	}

	// lakukan penghitungan rata-rata dan masukan ke channel ch
	ch <- float64(sum) / float64(len(num))

}

// func untuk mencari max value
func max(num []int, ch chan int) {
	// bikin variable untuk menampung nilai max
	var max = num[0]

	// looping element dan cek apakah nilai max lebih besar dari nilai yang ada di slice
	for _, val := range num {
		if val > max {
			max = val
		}
	}

	// masukkan nilai max ke channel ch
	ch <- max
}

func main() {

	// siapkan integer of slice
	nums := []int{1, 2, 3}
	fmt.Println("My Slice : ", nums)

	// sipakan channel
	ch1 := make(chan float64)
	ch2 := make(chan int)

	// isi parameter function dengan channel yang sudah disiapkan
	go average(nums, ch1)
	go max(nums, ch2)

	// lakukan selecting
	for i := 0; i < 2; i++ {
		select {
		case avg := <-ch1:
			fmt.Println("Average : ", avg)
		case max := <-ch2:
			fmt.Println("Max : ", max)
		default:
			fmt.Println("Nothing")
		}
	}

}
