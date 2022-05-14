package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// func AccessStudent
func AccessStudent(w http.ResponseWriter, r *http.Request) {

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

	// ====================== MUX DEFAULT ======================
	// // disini kita akan menggunakan mux sebagai penerapan dari inteface http.Handler

	// // inisiate mux
	// myMux := http.DefaultServeMux

	// // daftarkan url
	// myMux.HandleFunc("/student", AccessStudent)

	// // bikin handler untuk menginisialisasi middleware
	// var handler http.Handler = myMux
	// handler = MiddlewareAuth(handler)
	// handler = MiddlewareAllowOnlyGet(handler)

	// // bikin server
	// server := new(http.Server)
	// server.Addr = ":8080"
	// server.Handler = handler

	// // notif
	// fmt.Println("server on http://localhost:8080")
	// server.ListenAndServe()

	// ====================== MUX CUSTOM ======================
	// kita akan custom mux
	// inisiate mux
	myMux := new(CustomMux) //CustomMux ini akan kita bikin nanti di file middleware.go

	// bikin handlefunc sebagai route
	myMux.HandleFunc("/student", AccessStudent)

	// register middleware
	myMux.RegisterMiddleware(MiddlewareAuth)
	myMux.RegisterMiddleware(MiddlewareAllowOnlyGet)

	// create server
	server := new(http.Server)
	server.Addr = ":8080"
	server.Handler = myMux

	// notif
	fmt.Println("server on http://localhost:8080")
	server.ListenAndServe()
}
