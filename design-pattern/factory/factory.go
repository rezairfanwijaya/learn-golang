package factory

import "fmt"

/*
	Desain patern factory ini memiliki kesamaan seperti pewarisan pada bahasa berkonsep OOP. Jadi nanti kita akan memiliki sebuah class/object induk (super class) yang memiliki method yang nantinya wajib dipakai oleh turunannya (child class)

	Contoh kasus misalnya:
	Youtuber, nah youtuber ini kan pasti memilki konten yang mana konten ini bisa di play dan biasanya memiliki tipe konten (hiburan, edukasi, dan lain-lain). Dan youtuber ini bisa kita katakan sebagai content creator

	Nah jadi nanti kita akan memilki 2 super class yaitu Content yang akan memilki method play dan juga type lalu ada ContentCreator yang akan memiliki method CreateContent

	***** OK LETS CODE *****

*/

// super class content
type Content interface {
	Play()
	Type() string
}

// turunan dari content
type WebDevelopment struct{}
type MobileDevelopment struct{}
type VlogLiburan struct{}

var Edukasi = "Edukasi"
var Hiburan = "Hiburan"

// function untuk menjalankan content, masing-masing harus mengimplementasikan method dari interface content (super class)
func (w *WebDevelopment) Play() {
	fmt.Println("Hari ini sedang berlangsung konten web development")
}

func (w *WebDevelopment) Type() string {
	return Edukasi
}

func (m *MobileDevelopment) Play() {
	fmt.Println("Hari ini sedang berlangsung konten mobile development")
}

func (m *MobileDevelopment) Type() string {
	return Edukasi
}

func (v *VlogLiburan) Play() {
	fmt.Println("Hari ini sedang berlangsung konten vlog liburan")
}

func (v *VlogLiburan) Type() string {
	return Hiburan
}

// super class content creator
type ContentCreator interface {
	CreateContent() Content
}

// turunan dari content creator
type RezaIrfan struct{}
type Eko struct{}
type Atta struct{}

// kita implement method CreateContent dari interface content creator kepada masing masing struct turunan content creator
func (r *RezaIrfan) CreateContent() Content {
	return &WebDevelopment{}
}

func (e *Eko) CreateContent() Content {
	return &MobileDevelopment{}
}

func (a *Atta) CreateContent() Content {
	return &VlogLiburan{}
}
