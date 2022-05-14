package main

import (
	"log"
	"net/http"
	"time"

	gubrak "github.com/novalagung/gubrak/v2"
)

/*
	Cookie ini merupakan sebuah data dalam bentuk text yang disimpan di dalam browser ketika seseorang mengakses website.

	Cookie ini bisa berisikan berbagai macam informasi seperti, username, role dll.

	Kita akan coba membuat cookie sederhana
*/

// set nama cookie terlebih dahulu, ini akan muncul di Header
var CookieSaya = "MyCookie"

func main() {
	// buat handler
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/delete", DeleteHandler)

	log.Println("Server berjalan pada http://localhost:8080")
	// buat server
	http.ListenAndServe(":8080", nil)
}

// handler ini akan berfungsi untuk membuat cookie jika belum ada atau menimpa cookie baru jika cookie lama sudah expire
func IndexHandler(w http.ResponseWriter, r *http.Request) {

	// inisiasi object cookie
	c := &http.Cookie{}

	// kita cek apakah sudah ada cookie dengan nama CookieSaya sebelumnya
	if isExist, _ := r.Cookie(CookieSaya); isExist != nil {
		log.Printf("Cookie %s sudah ada", CookieSaya)
		// jika sudah ada, maka kita akan menimpa cookie lama dengan cookie baru
		c = isExist
		log.Printf("Cookie sebelum nya berhasil di replace dengan %s", c.Value)
	}

	// jika tidak ada maka kita akan buatkan cookie baru
	if c.Value == "" {
		c = &http.Cookie{
			Name:    CookieSaya,
			Value:   gubrak.RandomString(10), // value ini yang saya maksud bisa berisi string apapun terserah kita
			Expires: time.Now().Add(time.Minute * 10),
		}

		// simpan cookie
		http.SetCookie(w, c)
	}

	// tampilkan cookie
	w.Write([]byte(c.Value))
}

// functio ini akan menghapus cookie yang sudah dibuat tadi
func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	// cara menghapus cookie sangat mudah, cukup dengan deklatasi ulang cookie,dan mengisi nilai expiresnya dengan time.Unix(0,0) dan juga maxage -1
	c := &http.Cookie{
		Name:    CookieSaya,
		Expires: time.Unix(0, 0),
		MaxAge:  -1,
	}

	// set cookie
	http.SetCookie(w, c)
	w.Write([]byte("Cookie berhasil dihapus"))
}
