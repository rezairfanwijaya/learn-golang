package main

// Kenapa nama pakcage nya main ?
// karena digolang ada 2 jenis package
// 1. Executeable
// ini default dari golang nya yaitu package main -> ini bisa langsung di execute dari terminal

// 2. Library
// ini custom yang kita inginkan, jadi nama package nya selain main , contoh package keren, package naruto dlsb. package ini ga bisa langsung di execute dari teminal. harus masuk dulu ke program utama

// untuk import file dari package yang bebeda ada syntac nya
// yaitu modul/namapackage
// modul ini di dapat dari gile go.mod dan go.mod ini di dapat dari hasil inisiasi golang pas awal pertama membua project (go mod init namaProject)
// dan di porject ini nama module nya adalah myapp
import (
	"fmt"
	"myapp/calculation"
	"myapp/variables"
)

func main() {
	fmt.Println("Hallo Dunia")

	// memanggil funtion Sentence di file sentence.go
	// bisa langsung di panggil tanpa di import karena masih dalam satu package yaitu package main

	// bikin variable yang isi nya function Sentence
	kalimat := Sentence()
	fmt.Println(kalimat)

	// variable menampung hasil penjumlahan dari function Add di file add.go
	jumlah := calculation.Add(4, 90)
	fmt.Println("Hasil penjumlahan", jumlah)

	// variable menampung hasil perkalian dari function Mutiplication di file multiplication.go
	kali := calculation.Multipilcation(10, 90)
	fmt.Println("Hasil Perkalaian", kali)

	// variable untuk mencetak variable dari function kata di file kata.go
	kata := variables.Kata()
	fmt.Println(kata)

	// variable angka
	umur := 10
	fmt.Print(umur)

	// ++++++++++++++++ PERCABANGAN IF  ++++++++++++++++
	if umur > 15 {
		fmt.Println("Silahakn masuk ruangan")
	} else {
		fmt.Println("jangan masuk ruangan")
	}

	score := 78
	var index string

	index = "a"

	if score <= 50 {
		index = "C"
	} else if score > 60 && score < 80 {

		if score >= 75 && score <= 80 {
			index = "B+"
		} else {
			index = "B"
		}
	} else {
		index = "A"
	}

	fmt.Println(index)

	// ++++++++++++++++ SWITCH CASE  ++++++++++++++++
	size := 2
	switch size {
	case 1:
		fmt.Println("size", size)
	case 2:
		fmt.Println("size", size)
	case 3:
		fmt.Println("size", size)
	default:
		fmt.Println("size tidak dikenal")
	}
}
