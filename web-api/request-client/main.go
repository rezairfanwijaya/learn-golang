package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// siapkan struct untuk menampung data
type Student struct {
	ID    string
	Name  string
	Grade int
}

// base URL
var baseURL = "http://localhost:8080"

// function untuk mengirimkan request POST ke server
func fetchUsers() ([]Student, error) {
	// inisiasi
	var err error
	var client = &http.Client{}
	var data []Student

	// buat request
	request, err := http.NewRequest("GET", baseURL+"/users", nil)
	if err != nil {
		return nil, err
	}

	// lakukan aksi terhadap request yang telah dibikin
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	fmt.Println(response)

	// defer
	defer response.Body.Close()

	// decode hasil response (body/ data yang ditulis client)
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func main() {
	data, err := fetchUsers()
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println(data)
}
