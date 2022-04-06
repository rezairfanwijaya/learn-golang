package main

import (
	"fmt"
	"net/url"
)

func main() {
	// siapkan url
	var myUrl = "http://kalipare.com:80/hello?name=john wick&age=27"

	// panggil function parse
	u, err := url.Parse(myUrl)
	if err != nil {
		fmt.Println("Gagal memparsing url")
		return
	}

	fmt.Println("Url asal :", myUrl)

	// kita bisa mengambil protokol dari url
	protokol := u.Scheme
	fmt.Println("Protokol :", protokol)
	// kita bisa mengambil host dari url
	host := u.Host
	fmt.Println("Host :", host)
	// kita bisa mengambil path dari url
	path := u.Path
	fmt.Println("Path :", path)

	// kita juga bisa mengambil data yang di inputkan user
	fmt.Println("")
	// kita ambil nama
	nama := u.Query().Get("name")
	fmt.Println("Nama :", nama)
	// kita ambil usia
	usia := u.Query().Get("age")
	fmt.Println("Usia :", usia)

}
