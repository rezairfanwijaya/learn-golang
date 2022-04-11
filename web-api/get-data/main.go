package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// sipakan struct
type Student struct {
	ID    string
	Name  string
	Grade int
}

// isi struct Student
var Group = []Student{
	{"R123", "Reza", 3},
	{"R124", "Riza", 1},
	{"R125", "Adinda", 2},
}

// bikin handle function (controller) untuk menampilkan semua data Group
func Users(w http.ResponseWriter, r *http.Request) {
	// atur header
	w.Header().Set("Content-Type", "application/json")

	// cek method (GET, POST, PUT, DELETE)
	if r.Method == "GET" {
		// ubah struct Group diatas kedalam json
		byteGroup, err := json.Marshal(Group)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// tampilkan hasil
		w.Write(byteGroup)
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}

// bikin handlefunc(controller) untuk menampilkan data berdasarkan ID
func User(w http.ResponseWriter, r *http.Request) {
	// set header
	w.Header().Set("Content-Type", "application/json")

	// cek method (GET, POST, PUT, DELETE)
	if r.Method == "GET" {
		// ambil id yang di request client
		id := r.FormValue("id")
		var match []string

		// kita loop struct Group dan cocokan dengan Id yang di request client
		for _, val := range Group {
			if val.ID == id {
				// marshal (encode) struct ke dalam json
				jsonByte, err := json.Marshal(val)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				// add data ke array match
				match = append(match, string(jsonByte))
			}
		}

		// jika pangang array match kosong berarti tidak ada id yang sesuai
		if len(match) == 0 {
			http.Error(w, "Id not found", http.StatusNotFound)
			return
		}

		// tampilkan hasil
		w.Write([]byte(match[0]))

	}

	http.Error(w, "", http.StatusBadRequest)
}

func main() {
	http.HandleFunc("/users", Users)
	http.HandleFunc("/user", User)
	fmt.Println("server berjalan di http://localhost:8080")
	http.ListenAndServe(":8080", nil)

}
