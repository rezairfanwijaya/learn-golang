package main

import (
	"fmt"
	"regexp"
)

// regex berfungsi sebagai filtering dari data yang dikirimkan
// untuk filtering tersebut kita menggunakan function regex.Compile() dengan return regexp.*Regexp

func main() {
	// siapkan data
	var data string = "banana burger soup"
	fmt.Printf("\nKata yang akan difilter: %s\n", data)

	// panggil function (untuk menyiapkan filter yang akan dipakai)
	filterRegex, err := regexp.Compile(`[a-z]+`)
	if err != nil {
		panic(err)
	}

	// mulai filtering dengan (findAllString)--> mencari semua string yang sesuai dengan filter dan menampilkannya ke dalam array
	res1 := filterRegex.FindAllString(data, 2)
	res2 := filterRegex.FindAllString(data, 1)

	fmt.Println("res1: ", res1)
	fmt.Println("res2: ", res2)

	// maksud dari filter regex adalah mencari kata yang mengandung huruf a-z (lowercase)

	// dan findAllString() akan mengembalikan nilai array of string dengan parameter (kata yang dicari, jumlah kata yang dicari)

	// ===========================================

	// mulai filtering dengan (matchString)->mencocokan kata yang sesuai dengan filter dan mengembalikan boolean
	res3 := filterRegex.MatchString(data)
	fmt.Println("Apakah ada kata (true/false) :", res3)

	// ===========================================

	// mulai filtering dengan (findString)->mencocokan kata yang sesuai dengan filter dan mengembalikan nilai string
	cari := "KAMU"
	res4 := filterRegex.FindString(cari)
	fmt.Println("Kata yang dicari:", res4)

	// ===========================================

	// mulai filtering dengan (findStringIndex)->mencocokan kata yang sesuai dengan filter dan mengembalikan index dari array
	res5 := filterRegex.FindAllStringIndex(data, 1)
	fmt.Println("Index dari kata yang dicari:",
		res5)
	fmt.Println("kata tersebut:", data[res5[0][0]:res5[0][1]])

	// ===========================================

	// mulai filtering dengan (ReplaceAllString)->akan mereplace semua kata yang lolos proses filering
	res6 := filterRegex.ReplaceAllString(data, "ganti")
	fmt.Println("Setelah direplace", res6)
	fmt.Println("Sebelum direplace", data)

	// ===========================================

	// mulai filtering dengan (ReplaceAllStringFunc)->akan mereplace semua kata yang lolos proses filering dengan penambahan kondisi tertentu
	res7 := filterRegex.ReplaceAllStringFunc(data, func(s string) string {
		if s == "banana" {
			return "banana ganti"
		}
		return s
	})

	fmt.Println("Setelah direplace", res7)

	// ===========================================

	// mulai filtering dengan (Split)->akan memecah kata yang lolos proses filering menjadi array
	// kita siapkan filtering khusus untuk spilt ini dengan menggunakan function regex.Compile()

	// filter ini akan mengabaikan karakter yang tidak diinginkan (yaitu kata a dan b)
	split, _ := regexp.Compile(`[a-b]+`)
	res8 := split.Split(data, -1)
	fmt.Println("Setelah dipisah", res8)
}
