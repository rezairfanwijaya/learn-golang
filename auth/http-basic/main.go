package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// func AccessStudent
func AccessStudent(w http.ResponseWriter, r *http.Request) {

	// kita cek apakah user sudah melakukan auth
	if !Auth(w, r) {
		return
	}

	// kita cek apkakah method yang dikirim adalah GET? karena hanya GET yang diperbolehkan
	if !AllowOnlyGet(w, r) {
		return
	}

	// lalu kita cek lagi apakah di url ada parameter id jika ada maka tampilkan student berdasar id
	if id := r.URL.Query().Get("id"); id != "" {
		OutputJSON(w, SelectStudent(id))
		return
	}

	// Jika tidak maka tampilkan semua student
	OutputJSON(w, GetStudents())

}

// func outputJSON
func OutputJSON(w http.ResponseWriter, data interface{}) {
	res, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)

}

func main() {
	// daftarkan routing untuk akses student
	http.HandleFunc("/student", AccessStudent)

	// bikin server
	server := new(http.Server)
	server.Addr = ":8080"

	fmt.Println("Server is running at port 8080")
	server.ListenAndServe()

}
