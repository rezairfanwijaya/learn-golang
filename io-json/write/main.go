package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

// kita siapkan struct
type Student struct {
	Name  string `jsno:"name"`
	Score int    `json:"score"`
	Class string `json:"class"`
}

// kita bikin function untuk targeting path
func openFile(fileName string) (*os.File, error) {

	// ini jika file yang kita targeting sudah ada maka akan dibuka jika belum ada maka akan dibikinkan
	path, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic("Gagal membuka file")
	}

	return path, nil
}

// lali ini function untuk menulis data ke file json'
func writeData(fileName string, data []Student) error {

	// targeting path
	path, err := filepath.Abs(fileName + ".json")
	if err != nil {
		panic("Gagal membuka file")
	}

	// open file
	file, err := openFile(path)
	if err != nil {
		panic("Gagal membuka file")
	}

	// defer
	defer file.Close()

	// lalu kita tulis data struct ke file json
	jsonData, _ := json.Marshal(data)
	ioutil.WriteFile(path, jsonData, os.ModePerm)

	return nil
}

func main() {
	// masukan data ke struct
	data := []Student{
		{"Abdas", 90, "XI RPL 1"},
		{"Budi", 80, "XI RPL 2"},
		{"Caca", 70, "XI RPL 3"},
	}

	// filname
	fileName := "dataSiswa"

	writeData(fileName, data)
}
