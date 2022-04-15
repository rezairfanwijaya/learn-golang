package main

// import package database mysql
import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// bikin struct untuk menampung data hasil query
type Student struct {
	Id    int
	Name  string
	Age   int
	Grade int
}

// function untuk conncet ke database
func Connect_db() (*sql.DB, error) {

	// buat variable db
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/learn_go")
	// parameter sql.Open
	// 1. nama driver mysql
	// 2. username:password@tcp(host:port)/nama_database. Host bisa menggunakan 'localhost' / 127.0.0.1

	// handle error
	if err != nil {
		return nil, err
	}

	// return
	return db, err
}

// function untuk membaca data dari database
func Show_All_Data() {

	// panggil function connect_db
	db, err := Connect_db()
	if err != nil {
		fmt.Println("Gagal terhubung ke database")
		return
	}

	// defer
	defer db.Close()

	// query
	res, err := db.Query("SELECT * FROM tb_students")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	
	// defer
	defer res.Close()

	// bikin variable untuk menampung data hasil query
	var result []Student

	// looping
	for res.Next() {
		var student Student
		// scan data dari query ke struct
		err = res.Scan(&student.Id, &student.Name, &student.Age, &student.Grade)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		// append data ke variable result
		result = append(result, student)
	}

	// cek apakah database kosong
	if len(result) == 0 {
		fmt.Println("Database kosong")
		return
	}

	// looping hasil
	for _, v := range result {
		fmt.Println(v)
	}

}

// function untuk menampilkan data berdasarkan umur
func Show_Data_By_Age(age int) {

	// panggil function connect_db
	db, err := Connect_db()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// defer
	defer db.Close()

	// query
	res, err := db.Query("SELECT * FROM tb_students WHERE age = ?", age)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// defer
	defer res.Close()

	// bikin variable untuk menampung data hasil query
	var result []Student

	// looping
	for res.Next() {
		var student Student

		// scan data dari query ke struct
		res.Scan(&student.Id, &student.Name, &student.Age, &student.Grade)

		// append data ke variable result
		result = append(result, student)
	}

	// cek apakah ada data atau tidak
	if len(result) == 0 {
		fmt.Printf("Data dengan usia %d tidak ditemukan", age)
		return
	}

	// looping hasil
	fmt.Printf("\nMenampilkan hasil pencarian usia %d\n", age)
	for _, v := range result {
		fmt.Println(v)
	}

}

// function untuk menampilkan data berdasarkan id
// karena ID itu pasti unik, maka kita bisa menggunakan funtcion db.QueryRow() bukan db.Query()
func Show_Data_ById(id int) {

	// panggil function connect_db
	db, err := Connect_db()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// defer
	defer db.Close()

	// query
	// dan bisa dichain langsung dengan scan
	// sipakan varible penampung segaligus representasi data
	var res Student
	err = db.QueryRow("SELECT * FROM tb_students WHERE id = ?", id).Scan(&res.Id, &res.Name, &res.Age, &res.Grade)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// cek apakah ada data atau tidak
	if res.Id == 0 {
		fmt.Printf("Data dengan id %d tidak ditemukan", id)
		return
	}

	// tampilkan hasil
	fmt.Printf("\nMenampilkan hasil pencarian id %d\n", id)
	fmt.Println(res)

}

func main() {
	age := 29
	id := 2
	Show_All_Data()
	Show_Data_By_Age(age)
	Show_Data_ById(id)
}
