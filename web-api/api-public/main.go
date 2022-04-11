package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	MyUrl := "https://jsonplaceholder.typicode.com/users"
	resp, err := http.Get(MyUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	res := string(body)

	log.Println(res)

}
