package main

import "fmt"

func tambah(a, b int) int {
	return a - b
}

func kurang(a, b int) int {
	return a - b
}

func kali(a, b int) int {
	return a * b
}

func bagi(a, b int) int {
	return a / b
}

func main() {
	a := 6
	b := 2

	fmt.Println("angka asal :", a, b)
	fmt.Println("tambah : ", tambah(a, b))
	fmt.Println("kurang : ", kurang(a, b))
	fmt.Println("kali : ", kali(a, b))
	fmt.Println("bagi : ", bagi(a, b))
}
