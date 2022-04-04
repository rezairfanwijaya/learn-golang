package main

import (
	"encoding/csv"
	"os"
	"path/filepath"
)

// untuk menulis atau membuat file csv kita butuh 2 function seperti dibawah ini

// function ini akan membuka atau membuat file jiga ga ada
func openFile(csvName string) (*os.File, error) {

	// ambil direktori file csv
	// jika ada file nya maka ambil, jika ga ada maka bikin file tersebut
	path, err := filepath.Abs("../data/" + csvName + ".csv")
	if err != nil {
		return nil, err
	}

	// buka file csv atau bikin file jika belum ada (os.O_RDWR|os.O_CREATE)
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return nil, err
	}

	return file, nil
}

// function ini akan menyimpan kata yang kita inputkan ke file csv
func saveCSV(file string, records [][]string) error {

	// buka file menggunakan function openFile diatas
	f, err := openFile(file)
	if err != nil {
		return err
	}

	// close file pakai defer
	defer f.Close()

	// kita tulis file menggunakan package csv newWriter
	newRecord := csv.NewWriter(f)
	err = newRecord.WriteAll(records)
	if err != nil {
		return err
	}

	newRecord.Flush()
	return nil
}

func main() {
	// siapkan slice data untuk dimasukan ke file csv
	data := [][]string{
		{"nama", "umur", "alamat"},
		{"budi", "20", "jakarta"},
		{"andi", "30", "bandung"},
		{"siti", "40", "surabaya"},
	}

	// kita masukan data ke file csv melalui funtion saveCSV
	err := saveCSV("karyawanDB", data)
	if err != nil {
		panic(err)
	}

}
