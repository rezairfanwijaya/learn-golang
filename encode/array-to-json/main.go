package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// siapkan array
	var Mhs = []string{
		"Rizki Adam Kurniawan",
		"15.01.53.0012",
		"Jl. Kedung Baruk No.1",
		"TI-2A",
	}

	// lakukan marshal array ke json
	jsonMhs, _ := json.Marshal(Mhs)

	// konversi byte jsonMhs ke string
	jsonMhsFinal := string(jsonMhs)

	fmt.Println(Mhs)
	fmt.Println(jsonMhs)
	fmt.Println(jsonMhsFinal)
}
