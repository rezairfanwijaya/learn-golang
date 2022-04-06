package main

import (
	"fmt"
	"net/http"
	"text/template"
)

// golang web server ini akan menggunakan package net/http yang banyak sekali kegunaan.
// package net/http berisi fungsi-fungsi yang berguna untuk mengirimkan respon HTTP ke client.
// pada kali ini net/http akan kita gunakan untuk menjalankan server dan juga memunculkan text di web ketika sebuah url diakses

func hello(w http.ResponseWriter, r *http.Request) {
	// kita akan memunculkan text hello world di web ketika func hello ini dipanggil
	fmt.Fprintln(w, "Hello World")
}

func index(w http.ResponseWriter, r *http.Request) {
	// siapkan data yang akan dimunculkan ke halaman web
	var data = map[string]string{
		"nama":  "Reza Irfan Wijaya",
		"title": "GO web-server",
	}

	// kita panggil halaman html kita
	myTemplate, err := template.ParseFiles("index.html")
	if err != nil {
		fmt.Println("gagal memanggil template")
		return
	}

	// kita masukan data ke template
	myTemplate.Execute(w, data)

}

func main() {
	// buat route dan memanggil fungsi hello
	http.HandleFunc("/", hello)

	// buat route untuk menjalankan html template sederhana
	http.HandleFunc("/index", index)

	fmt.Println("Server berjalan pada : http://localhost:8080")

	// start server
	http.ListenAndServe(":8080", nil)
}
