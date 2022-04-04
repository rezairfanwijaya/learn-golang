package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// kita bikin function untuk update file
// function ini memiliki parameter berupa kata (string) yang akan ditulis

func updateFile(newWord string) {

	// kita siapkan file yang akan kita baca
	fileName := "nama.txt"

	// kita baca text dari file
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	// kita tampilkan data yang ada di file sebelum diupdate
	// untuk melihat perbandingan antara data sebelumnya dengan data yang baru
	fmt.Printf("\nBefore Update \n")
	fmt.Printf("Nama file :%s\n", fileName)
	fmt.Printf("Panjang data : %d\n", len(data))
	fmt.Printf("Data : %s\n\n", data)

	// proses update
	// kita pisahkan tiap kata di file nama.txt berdasarkan spasi
	lines := strings.Split(string(data), " ") // output array

	// fmt.Println(lines)

	// lakukan pengecekan tiap kata mengguakan perulangan
	// disini akan mencari kata "Nama" dan akan diganti menjadi kata baru sesuai dengan parameter function ini
	for i, line := range lines {
		// cek apakah ada kata "Nama"
		if strings.Contains(line, "Ok") {
			lines[i] = newWord
		}
	}

	// lalu tampung data hasil update tadi ke dalam variable output
	output := strings.Join(lines, " ")
	// fmt.Println(output)
	err = ioutil.WriteFile(fileName, []byte(output), 0644)
	if err != nil {
		panic(err)
	}

	// kita baca kembali file nama.txt setelah diupdate
	data, err = ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	// tampilkan data yang ada di file setelah diupdate
	fmt.Printf("\nAfter Update \n")
	fmt.Printf("Nama file :%s\n", fileName)
	fmt.Printf("Panjang data : %d\n", len(data))
	fmt.Printf("Data : %s\n\n", data)

}

func main() {
	updateFile("Nice")
	reset("Ok")
}

// kita juga bisa bikin function untuk mereset update yang kita lakukan

func reset(oldWord string) {
	//deklarasi file yang akan direset
	fileName := "nama.txt"

	// ambil isi file
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	// kita split data
	lines := strings.Split(string(data), " ")

	// cek tiap kata
	for i, line := range lines {
		if strings.Contains(line, "Nice") {
			lines[i] = oldWord
		}
	}

	// masukan hasil reset ke dalam variable output
	output := strings.Join(lines, " ")
	err = ioutil.WriteFile(fileName, []byte(output), 0644)
	if err != nil {
		panic(err)
	}

	// baca kembali file nama.txt
	data, err = ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	// tampilkan data
	fmt.Printf("\nAfter Reset \n")
	fmt.Printf("Nama file :%s\n", fileName)
	fmt.Printf("Panjang data : %d\n", len(data))
	fmt.Printf("Data : %s\n\n", data)

}
