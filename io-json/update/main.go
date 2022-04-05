package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// sipakan struct
type Teacher struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	NIP  string `json:"nip"`
}

// open file / create file
func openFile(fileName string) (*os.File, error) {

	path, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic("Path tidak ditemukan")
	}
	return path, nil
}

// write data
func writeJson(fileName string, data []Teacher) error {

	path, err := filepath.Abs(fileName + ".json")
	if err != nil {
		panic("Gagal membuka file")
	}

	file, err := openFile(path)
	if err != nil {
		panic("Gagal membuka file")
	}

	defer file.Close()

	jsonData, _ := json.Marshal(data)
	ioutil.WriteFile(path, jsonData, os.ModePerm)
	return nil
}

func readJson(fileName string) ([]Teacher, error) {
	path, err := filepath.Abs(fileName + ".json")
	if err != nil {
		panic("Gagal membuka file")
	}

	file, err := openFile(path)
	if err != nil {
		panic("Gagal membuka file")
	}

	defer file.Close()

	teacher := []Teacher{}

	byteData, _ := ioutil.ReadAll(file)
	json.Unmarshal(byteData, &teacher)
	return teacher, nil
}

// update data
func updateJson(fileName string, newData []Teacher) ([]Teacher, error) {
	path, err := filepath.Abs(fileName + ".json")
	if err != nil {
		panic("Gagal membuka file")
	}

	file, err := openFile(path)
	if err != nil {
		panic("Gagal membuka file")
	}

	defer file.Close()

	// proses update
	teacher, err := readJson(fileName)
	if err != nil {
		panic("Gagal membuka file")
	}
	// append
	teacher = append(teacher, newData...)

	// rewrite
	jsonData, _ := json.Marshal(teacher)
	ioutil.WriteFile(path, jsonData, os.ModePerm)
	return teacher, nil
}

func main() {
	fileName := "teacher"
	// oldData := []Teacher{
	// 	{"Budi", 80, "NIP 1"},
	// 	{"Caca", 70, "NIP 2"},
	// }

	newData := []Teacher{
		{"Budi", 80, "NIP 11"},
		{"Caca", 70, "NIP 21"},
	}

	// err := writeJson(fileName, oldData)
	// if err != nil {
	// 	panic("Gagal menulis file")
	// }

	data, err := updateJson(fileName, newData)
	if err != nil {
		panic("Gagal menulis file")
	}

	for _, v := range data {
		fmt.Println(v)
	}

	fmt.Println("Update Berhasil !!")

}
