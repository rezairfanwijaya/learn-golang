package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// siapkan struct untuk menampung data hasil decode json
type Siswa struct {
	Nama  string `json:"nama"`
	Nilai int    `json:"nilai"`
	Kelas string `json:"kelas"`
}

// bikin custom i/o stringer
func (s Siswa) String() string {
	return fmt.Sprintf("Nama : %s, Nilai : %d, Kelas : %s", s.Nama, s.Nilai, s.Kelas)
}

// kita bikin function untuk open file json nya
// function ini akan mentarget path dari file json yang akan dibaca
func openFile(path string) (*os.File, error) {
	file, err := os.OpenFile(path, os.O_RDWR, os.ModePerm)
	if err != nil {
		fmt.Println("Gagal membuka error")
	}
	return file, nil
}

// lalu kita bikin function untuk membaca file json yang sudah kita buka tadi path nya
func readData(fileName string) ([]Siswa, error) {

	// targerting path
	path, err := filepath.Abs(fileName + ".json")
	if err != nil {
		fmt.Println("Gagal membuka file")
	}

	// open file
	file, err := openFile(path)
	if err != nil {
		panic("Gagal membuka file")
	}

	// defer
	defer file.Close()

	// bikin slice untuk menampung data hasil decode json
	siswa := []Siswa{}

	// baca file json
	bytevalue, _ := ioutil.ReadAll(file)
	json.Unmarshal(bytevalue, &siswa)

	return siswa, nil
}

func main() {
	fileName := "siswa"
	data, err := readData(fileName)
	if err != nil {
		panic("Gagal membaca file")
	}
	for _, item := range data {
		fmt.Println(item)
	}
}
