package main

import (
	"log"
	"net/http"
)

func main() {
	// inisiasi server
	server := http.Server{
		Addr: ":8080",
	}

	// register handler dan route
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	// run server
	server.ListenAndServe()
}

// function handler route /
func Home(w http.ResponseWriter, r *http.Request) {
	// cek method
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		log.Fatalf("Method not allowed: %s", r.Method)
		return
	}

	// tampilkan reponse sukses
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Welcome to my Home"))
}

// function handler route /about
func About(w http.ResponseWriter, r *http.Request) {
	// cek method
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		log.Fatalf("Method not allowed: %s", r.Method)
		return
	}

	// tampilkan reponse sukses
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("This is About me"))
}
