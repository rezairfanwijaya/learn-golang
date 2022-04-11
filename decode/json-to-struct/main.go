package main

import (
	"encoding/json"
	"fmt"
)

// kita akan ubah json ke dalam bentuk struct

// siapkan struct telebih dahulu
type Mahasiswa struct {
	Nama   string `json:"nama"`
	Nim    string `json:"nim"`
	Alamat string `json:"alamat"`
	Kelas  string `json:"kelas"`
}

func main() {
	// siapkan json
	jsonString := `{"nama":"Rizki Adam Kurniawan","nim":"15.01.53.0012","alamat":"Jl. Kedung Baruk No.1","kelas":"TI-2A"}`

	// ubah json kedalam byte
	jsonByte := []byte(jsonString)

	// inisiasi struct
	var mhs Mahasiswa

	// decode dan masukan json ke dalam struct
	err := json.Unmarshal(jsonByte, &mhs)
	if err != nil {
		fmt.Println(err.Error())
	}

	// tampilkan hasil
	fmt.Println(mhs.Nama)
	fmt.Println(mhs.Alamat)
	fmt.Println(mhs.Kelas)
	fmt.Println(mhs.Nim)

}
