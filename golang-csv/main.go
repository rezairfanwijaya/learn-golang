package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
)

// file csv biasanya digunakan pada aplikasi excel, spreadsheet, atau database.
// jadi ini sangat berguna dan penting

// funtion untuk membuka file csv
func openFile(csvName string) (*os.File, error) {

	// ambil direktori file csv
	path, err := filepath.Abs("data/" + csvName + ".csv")
	if err != nil {
		return nil, err
	}

	// buka file csv
	file, err := os.OpenFile(path, os.O_RDWR, os.ModePerm)
	if err != nil {
		return nil, err
	}

	return file, nil
}

// function untuk membaca file csv
func readfile(file string) ([][]string, error) {
	// buka file menggunakan function openFile diatas
	f, err := openFile(file)
	if err != nil {
		return nil, err
	}

	// close file pakai defer
	defer f.Close()

	// kita baca file menggunakan package csv newReader
	newReader := csv.NewReader(f)
	records, err := newReader.ReadAll()
	if err != nil {
		return nil, err
	}

	return records, nil

}

func main() {
	// kita panggil function readfile
	data, err := readfile("siswaDb")
	if err != nil {
		panic(err)
	}

	// fmt.Println(data)
	for index, value := range data {
		// kita cek apakah index 0 atau 1
		if index == 0 {
			// kita cetak header
			fmt.Println(value)
		} else {
			fmt.Println(value)
		}
	}
}
