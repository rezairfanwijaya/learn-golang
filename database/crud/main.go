package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() (*sql.DB, error) {
	con, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/learn_go")
	if err != nil {
		return nil, err
	}
	return con, err
}

func CRUD() {

	// connect ke database
	conn, err := Connect()
	if err != nil {
		fmt.Println(err.Error())
	}

	// defer
	defer conn.Close()
	

	// insert data
	_, err = conn.Exec("INSERT INTO tb_students values (?, ?, ?, ?) ", 6, "EJA", 20, 1)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Data berhasil ditambahkan")
	}

	// update data
	_, err = conn.Exec("UPDATE tb_students SET name = ?, age = ?, grade = ? WHERE name = ?", "Budi Edit", 20, 3, "Budi")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Data berhasil diupdate")
	}

	// delete data
	_, err = conn.Exec("DELETE FROM tb_students WHERE id = ?", 2)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Data berhasil dihapus")
	}

}

func main() {
	CRUD()
}
