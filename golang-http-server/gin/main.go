package main

import (
	"github.com/gin-gonic/gin"
)

type Book struct {
	Judul     string `json:"judul"`
	Penerbit  string `json:"penerbit"`
	Pengarang string `json:"pengarang"`
	Tahun     string `json:"tahun"`
}

func main() {
	// inisiasi gin
	r := gin.Default()
	// daftarkan route dan handler
	r.GET("/", home)
	r.GET("/about", about)
	r.POST("/books", postBook)

	// run server
	r.Run(":8080")
}

func home(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Wellcome to my home",
	})
}

func about(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Wellcome to my about",
	})
}

func postBook(c *gin.Context) {

	// inisiasi struct book
	var myBook Book

	// ambil data yang dimasukan oleh client memakai souldbindjson
	err := c.ShouldBindJSON(&myBook)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Gagal Updload Buku",
		})
		return
	}

	// jika berhasil tampilakan data yang dimasukan
	c.JSON(200, gin.H{
		"message":   "Berhasil Upload Buku",
		"judul":     myBook.Judul,
		"penerbit":  myBook.Penerbit,
		"pengarang": myBook.Pengarang,
		"tahun":     myBook.Tahun,
	})
}
