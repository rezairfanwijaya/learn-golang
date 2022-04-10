package main

import "fmt"

// testify digunakan sebagai pengganti if else dalam pengecekan suatu program
// sehingga kita tidak perlu menuliskan kode if else lagi
// testify menggunakan kata kunci assert sebagai pengganti if else

func Sum(a, b int) (res int) {
	return a + b
}

func Hello(nama string) string {
	return fmt.Sprintf("Hello %s", nama)
}

func bagi(bil1, bil2 int) int {
	return bil1 / bil2
}

func main() {
	fmt.Println(Sum(4, 6))
	fmt.Println(Hello("Budi"))
	fmt.Println(bagi(10, 2))
}
