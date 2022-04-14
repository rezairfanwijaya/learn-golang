package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	
)

type Student struct {
	ID    string
	Name  string
	Grade int
}

var baseURL = "http://localhost:8080"

func fetchUser(ID string) (Student, error) {
	// inisiasi
	var err error
	var client = &http.Client{}
	var data Student

	// karena ini akan mengambil data berdasarkan ID , maka kita harus mengambil ID tersebut
	var param = url.Values{}
	param.Set("id", ID)
	var playload = bytes.NewBufferString(param.Encode())
	fmt.Println(playload)

	// buat request
	req, err := http.NewRequest("GET", baseURL+"/user", playload)
	if err != nil {
		return data, err
	}


	// set header
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// lakukan aksi terhadap request yang telah dibikin
	response, err := client.Do(req)
	if err != nil {
		return data, err
	}
	
	fmt.Println(response)

	// defer
	defer response.Body.Close()

	// decode data json
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return data, err
	}

	return data, nil

}

func main() {
	data, err := fetchUser("R124")
	if err != nil {
		log.Fatal(err)
		// fmt.Println("Gagal melakukan request")
		return
	}

	fmt.Printf("ID: %s Name: %s Grade: %d\n", data.ID, data.Name, data.Grade)
}
