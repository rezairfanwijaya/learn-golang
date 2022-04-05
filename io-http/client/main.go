package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type user struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type message struct {
	Data string `json:"message"`
}

var baseURL = "https://localhost:8080/"

// check server
func checkServer(user user) (message, error) {
	var err error
	var client = &http.Client{}
	userByte, _ := json.Marshal(user)
	var receivedMessage message
	// user melakuan persiapan request dengan params post ke server dengan user baseURL
	req, err := http.NewRequest("POST", baseURL, bytes.NewBuffer(userByte))
	if err != nil {
		return message{}, err
	}

	// set header
	req.Header.Set("Content-Type", "application/json")

	// user melalukan request ke server
	response, err := client.Do(req)
	if err != nil {
		return message{}, err
	}

	// defer
	defer response.Body.Close()

	// baca semua body yang di tulis user
	rbody, _ := ioutil.ReadAll(response.Body)

	// lalu buka json untuk dibaca
	err = json.Unmarshal(rbody, &receivedMessage)
	if err != nil {
		return message{}, err
	}

	return receivedMessage, nil

}

func main() {
	// kita coba melakukan input data dari user
	var name string
	var age int
	fmt.Printf("\ninput name : ")
	fmt.Scan(&name)
	fmt.Printf("\ninput age : ")
	fmt.Scan(&age)

	// lalu masukan inputan user ke stuct
	var newUser = user{
		Name: name,
		Age:  age,
	}

	// lalu kita coba melakukan request ke server
	fmt.Printf("i send this data : %v\n", newUser)
	revMessage, _ := checkServer(newUser)
	fmt.Printf("i receive this data : %v", revMessage)

}
