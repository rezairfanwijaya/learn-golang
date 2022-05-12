package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// url
var URL = "http://localhost:8080"

// inisiasi client
var client = &http.Client{}

func main() {
	home()
	about()

}

func home() {
	// client melakukan request get ke server melalui url
	resp, err := client.Get(URL)
	if err != nil {
		log.Fatal(err)
	}

	// defer resp agar tidak memakan banyak memori
	defer resp.Body.Close()

	// baca semua data dari body resp
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read body: %v", err)
	}

	// tampilkan data
	fmt.Println(string(data))
}

func about() {
	// client melakukan request get ke server melalui url + /about
	resp, err := client.Get(URL + "/about")
	if err != nil {
		log.Fatal(err)
	}

	// defer untuk menghemat memori
	defer resp.Body.Close()

	// baca semua data dari body resp
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read body: %v", err)
	}

	// tampilkan data
	fmt.Println(string(data))
}
