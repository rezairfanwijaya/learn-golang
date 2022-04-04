// disini kita akan mengimplementasikan hasil belajar kita tentang read-file , write-file dan update-file
// dengan membuat program menghitung jumlah angka yang ada di dalam file

package main

import (
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

// kita bikin function untuk membaca dan mengambil nilai yg ada fi file input.txt dengan kembalian byte

func readFile() []byte {

	// deklarasi file yang akan dibaca
	fileName := "input.txt"

	// ambil isi file
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	// return nilai data
	return data
}

// bikin function yang akan membuat file output.txt untuk menampung hasil dari penjumlahan
func createFile(result int) {

	// bikin file output.txt
	data, err := os.Create("output.txt")
	if err != nil {
		panic(err)
	}

	// tutup file
	defer data.Close()

	// kita konversi nilai result ke string
	strNumber := strconv.Itoa(result)

	// kita masukan data hasil penjumlahan ke dalam file output.txt
	_, err = data.WriteString(strNumber)
	if err != nil {
		panic(err)
	}

}

// kita bikin function untuk menjumlahkan data
func sum(input []byte) int {
	result := 0

	// kita harus konversi data byte ke interger
	// kita bikin dulu funtion untuk konversi (ByteToInt)
	number, err := BytetoInt(input)
	if err != nil {
		panic(err)
	}

	// lakukan penjumlahan
	for _, val := range number {
		result += val
	}

	return result

}

// function untuk konversi
func BytetoInt(input []byte) ([]int, error) {

	// kita split input nya
	byteString := strings.Split(string(input), " ")

	// sipkan variable slice untuk hasil konversi
	result := []int{}

	// lakukan konversi dari string hasil split ke interger
	for _, val := range byteString {
		integerFull, _ := strconv.Atoi(val)
		// add hasil ke dalam slice
		result = append(result, integerFull)
	}

	// return nilai result
	return result, nil
}

func main() {
	input := readFile()
	result := sum(input)
	createFile(result)

}
