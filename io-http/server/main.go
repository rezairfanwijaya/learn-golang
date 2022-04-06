package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type user struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type message struct {
	Data string `json:"message"`
}

func hello(w http.ResponseWriter, r *http.Request) {
	// set header
	w.Header().Set("Content-Type", "application/json")

	// set response
	if r.Method == "POST" {
		// kita baca semua data yang dikirim user dari body
		reqbody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// bikin varible penampung data user
		var newUser user

		// kita buka json untuk dibaca (unmarshal)->decode
		json.Unmarshal(reqbody, &newUser)
		log.Printf("i receive this data : %+v", newUser)

		// bikin variable untuk menampung message
		var msg message
		msg.Data = fmt.Sprintf("Hello %s, you are %d years old", newUser.Name, newUser.Age)
		log.Printf("i send this data : %+v", msg)

		// lalu kita encode ke json
		resultByte, _ := json.Marshal(msg)
		w.WriteHeader(http.StatusAccepted)
		w.Write(resultByte)
		return

	}
	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}

func main() {
	http.HandleFunc("/", hello)
	fmt.Println("server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
