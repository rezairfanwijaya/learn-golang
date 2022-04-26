package main

// import ini bisa diambil setelah mendownload third party bernama gubrak dari link berikut github.com/novalagung/gubrak/v2 dengan cara  go get -u github.com/novalagung/gubrak/v2
import (
	"fmt"

	"github.com/novalagung/gubrak/v2"
)

func main() {
	// ini akan generate int random 10-50
	rand := gubrak.RandomInt(10, 50)
	fmt.Println(rand)
}

// setelah func main dibuat lalu jalankan perintah pada terminal dan arahakn path ke file main.go yang menggunakan 3rd party gubrak

// go mod vendor -> untuk melakukan vendoring 3rd party library yang digunakan
// setelah berhasil dieksekusi maka akan muncul folder vendor pada root

// go build -o myvendor.exe -> untuk mengeksekusi file main.go yang menggunakan 3rd party gubrak dan menghasilkan file executable

// manfaat vendoring ini sama seperti kita menggunakan bootstrap dengan cdn atau download langsung menggunakan NPM
// nah jika menggunakan cdn (tidak vendoring) ini kita akan tergantung terhadap link yang kita gunakan
// tapi jika menggunakan vendoring ini kita tidak tergantung terhadap link yang kita gunakan dan jika ada update maka program kita tidak akan terpengaruh
