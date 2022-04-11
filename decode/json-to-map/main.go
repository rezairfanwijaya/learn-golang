package main

import (
	// "encoding/json"
	"encoding/json"
	"fmt"
)

func main() {

	// sipakan json
	jsonString := `{
		"nama": "Rizki Adam Kurniawan",
		"nim": "15.01.53.0012",
		"alamat": "Jl. Kedung Baruk No.1",
		"kelas": "TI-2A"
	}`

	// siapkan map
	var map1 map[string]interface{}
	json.Unmarshal([]byte(jsonString), &map1)

	for _, val := range map1 {
		fmt.Println(val)
	}

	fmt.Println(map1["nama"])
}
