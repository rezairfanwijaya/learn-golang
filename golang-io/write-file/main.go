package main

import (
	"fmt"
	"os"
)

// kita bikin function untuk membuat file
func createFile() {

	// bikin nama file
	file, err := os.Create("write.txt")
	if err != nil {
		panic(err)
	}

	// kita tutup file nya setelah kita baca
	defer file.Close()

	// kita isi teks kedalam file yang telah kita bikin
	len, err := file.WriteString("Hello World\n" + "Saya sedang belajar golang")
	if err != nil {
		panic(err)
	}

	// kita debug
	fmt.Printf("Nama file : %s", file.Name())
	fmt.Printf("\nPanjang file : %d", len)
}

func main() {

	// panggil function createFile
	createFile()
}
