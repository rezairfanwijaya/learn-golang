package main

import (
	"fmt"
	"io/ioutil"
)

// function untuk membaca file
func readFile() {

	// nama file yang akan dibaca
	fileName := "read.txt"

	// lalu kita baca file nya
	result, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	// kita cetak hasil nya
	fmt.Printf("Nama file : %s\n", fileName)
	fmt.Printf("Isi file : %s\n", result)
	fmt.Printf("Panjang file : %d\n", len(result))

}

func main() {
	// panggil function readFile
	readFile()
}
