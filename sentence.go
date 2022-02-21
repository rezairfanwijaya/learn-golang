package main

import "fmt"

// function akan mengembalikan string , maka disamping nama function di tulis string

// nama funtion (!=main) harus diawali huruf besar
func Sentence() string {
	// Sprintf itu untuk return string dia hanya membawa , tidak mencetak. untuk cetak tetap harus Println
	return fmt.Sprintf("Hallo saya %s saya sedang belajar %s", "Jhon", "Golang")
}
