package main

// Kenapa nama package nya main ?
// karena digolang ada 2 jenis package
// 1. Executeable
// ini default dari golang nya yaitu package main -> ini bisa langsung di execute dari terminal

// 2. Library
// ini custom yang kita inginkan, jadi nama package nya selain main , contoh package keren, package naruto dlsb. package ini ga bisa langsung di execute dari teminal. harus masuk dulu ke program utama

// untuk import file dari package yang bebeda ada syntac nya
// yaitu modul/namapackage
// modul ini di dapat dari gile go.mod dan go.mod ini di dapat dari hasil inisiasi golang pas awal pertama membua project (go mod init namaProject)
// dan di porject ini nama module nya adalah myapp
import (
	"errors"
	"fmt"
	"myapp/buku"
	"myapp/calculation"
	"myapp/variables"
)

// ++++++++++++++++ STRUCT  ++++++++++++++++
// deklarasi type namaStuct struct
type User struct {
	// deklarasi variable atau properties
	ID        int
	FirstName string
	LastName  string
	Email     string
	IsActive  bool
}

// ++++++++++++++++ EMBEDDED STRUCT  ++++++++++++
// adalah stuct yang memilki properti berupa struct lainnya
// case kali ini saya akan mendemokan sebuah group mahasiswa di telegram yang mana group tersebut akan memiliki nama, admin, dan anggota

// buat struct anggota
type mahasiswa struct {
	Nama     string
	Nim      int
	Fakultas string
	Jurusan  string
	Kelas    string
	Usia     int
}

// kedua buat struct untuk group telegram
type group struct {
	Nama string
	// admin ini bertipe mahasiswa karena, properti yg dimiliki admin akan sama dengan properti yang dimiliki struct mahasiswa
	Admin mahasiswa
	// sama dengan admin, anggota group pun akan memilki properti yang sama dengan struct mahasiswa dan kita bikin dia jadi slice karena anggota bisa lebih dari satu, beda dengan admin yang hanya 1
	Anggota     []mahasiswa
	isAvailable bool
}

// ++++++++++++++++ METHODE STRUCT  ++++++++++++
// method pada struct ini hampir mirip seperti function pada umumnya namun memiliki syntax yang berdeda
// disini saya akan mencoba membuat struct baru berupa data negara sekaligus membuat method untuk mencetak data negara tersebut

type Nation struct {
	Nama     string
	Penduduk int64
	Luas     int64
	Presiden string
	Benua    string
}

// nah jadi syntax method itu seperti ini
func (negara Nation) show() string {
	return fmt.Sprintf("Nama Negara : %s, jumlah penduduk: %d", negara.Nama, negara.Penduduk)
}

// ++++++++++++++++ QUIZ STRUCT  ++++++++++++
// pada quiz ini saya akan mengambil case struct untuk organisasi osis yang terdiri dari beberapa siswa
type Siswa struct {
	Nama     string
	Kelas    string
	Angkatan int
	Motivasi string
}

type Osis struct {
	Nama      string
	Ketua     string
	Wakil     string
	Bendahara string
	Anggota   []Siswa
}

func (group Osis) showOsis() {
	fmt.Print("\n\n")
	fmt.Printf("OSIS %s\n\n", group.Nama)
	fmt.Printf("Ketua     :  %s\n", group.Ketua)
	fmt.Printf("Wakil     :  %s\n", group.Wakil)
	fmt.Printf("Bendahara :  %s\n\n", group.Bendahara)
	fmt.Println("Daftar Anggota : ")
	for _, anggota := range group.Anggota {
		fmt.Printf("Nama    : %s\n", anggota.Nama)
		fmt.Printf("Kelas 	: %s\n", anggota.Kelas)
		fmt.Printf("Angkatan: %d\n", anggota.Angkatan)
		fmt.Printf("Motivasi: %s\n\n", anggota.Motivasi)
	}
}

func main() {

	fmt.Println("Hallo Dunia")

	// memanggil funtion Sentence di file sentence.go
	// bisa langsung di panggil tanpa di import karena masih dalam satu package yaitu package main

	// bikin variable yang isi nya function Sentence
	kalimat := Sentence()
	fmt.Println(kalimat)

	// variable menampung hasil penjumlahan dari function Add di file add.go
	jumlah := calculation.Add(4, 90)
	fmt.Println("Hasil penjumlahan", jumlah)

	// variable menampung hasil perkalian dari function Mutiplication di file multiplication.go
	kali := calculation.Multipilcation(10, 90)
	fmt.Println("Hasil Perkalaian", kali)

	// variable untuk mencetak variable dari function kata di file kata.go
	kata := variables.Kata()
	fmt.Println(kata)

	// variable angka
	umur := 10
	fmt.Print(umur)

	// ++++++++++++++++ PERCABANGAN IF  ++++++++++++++++
	if umur > 15 {
		fmt.Println("Silahakn masuk ruangan")
	} else {
		fmt.Println("jangan masuk ruangan")
	}

	score := 78
	var index string

	index = "a"

	if score <= 50 {
		index = "C"
	} else if score > 60 && score < 80 {

		if score >= 75 && score <= 80 {
			index = "B+"
		} else {
			index = "B"
		}
	} else {
		index = "A"
	}

	fmt.Println(index)

	// ++++++++++++++++ SWITCH CASE  ++++++++++++++++
	size := 2
	switch size {
	case 1:
		fmt.Println("size", size)
	case 2:
		fmt.Println("size", size)
	case 3:
		fmt.Println("size", size)
	default:
		fmt.Println("size tidak dikenal")
	}

	// ++++++++++++++++ LOOPING  ++++++++++++++++
	// for
	iter := 100
	for i := 1; i <= iter; i++ {
		fmt.Println(i, "saya siapa ? ")
	}

	// while nya golang
	i := 1
	for i <= 10 {
		fmt.Println("saya dari while", i)
		i++
	}

	// foreach nya golang
	data := "Reza Irfan wijaya"
	for index, huruf := range data {
		fmt.Println("index :", index, "hurufnya : ", string(huruf))
	}

	// lalu jika ingin menampilkan huruf saja tinggal dengan cari berikut
	for _, letters := range data {
		fmt.Println("Huruf nya : ", string(letters))
	}
	// syntax _ digunakan untuk menampung/memperbolehkan varibale yang tidak digunakan untuk tetap dideklarasikan

	// saya akan mencoba latihan
	// ada string GOLANG ADALAH BAHASA TERBAIK
	// jika index ganjil jangan tampilkan
	golang := "GOLANG BAHASA TERBAIK SAAT INI"
	for i, huruf := range golang {

		// seleksi index ganjil
		if i%2 == 0 {

			fmt.Println("huruf hasil seleksi ganjil genap: ", string(huruf))

			// menggunakan if
			// seleksi huruf vokal
			if string(huruf) == "A" || string(huruf) == "I" || string(huruf) == "U" || string(huruf) == "E" || string(huruf) == "O" {
				fmt.Println("Huruf hasil seleksi vokal", string(huruf))
			}

			// menggunakan case
			switch string(huruf) {
			case "A", "I", "U", "E", "O":
				fmt.Println("huruf VOKAL dari case: ", string(huruf))

			}
		}

	}

	// ++++++++++++++++ ARRAY  ++++++++++++++++
	// deklarasi cara ke 1
	var mobils [5]string
	mobils[0] = "Lambo"
	mobils[1] = "Ferrari"
	mobils[2] = "BMW"
	mobils[3] = "Mazda"
	mobils[4] = "Angkot"
	// cetak semua isi
	fmt.Println(mobils)
	// cetak data tertentu dari array
	fmt.Println(mobils[0])
	fmt.Println(mobils[4])
	// cetak panjang array
	fmt.Println(len(mobils))

	// deklarasi cara ke 2
	// motors := [4]string{"ZX10R", "HAYABUSA", "VESPA", "HARLEY"}
	// bisa juga isi vertikal, namun data terakhir harus ada koma (Setelah harley)
	// motors := [4]string{
	// 	"ZX10R",
	// 	"HAYABUSA",
	// 	"VESPA",
	// 	"HARLEY",
	// }

	// selain itu array juga bisa menentukan panjang nya sendiri dengan menggunakan [...] misal
	motors := [...]string{
		"ZX10R",
		"HAYABUSA",
		"VESPA",
		"HARLEY",
		"JUPITER",
		"SUPRA",
	}
	// cetak semua isi
	fmt.Println(motors)
	// cetak data tertentu dari array
	fmt.Println(motors[3])
	fmt.Println(motors[1])
	// cetak panjang array
	fmt.Println(len(motors))

	// perulangan pada array
	for index, motor := range motors {
		fmt.Println("Motor dengan index ke -", index, motor)
	}

	// ++++++++++++++++ SLICE  ++++++++++++++++
	// deklarasi pertama
	var guitars []string
	// memasukan data ke array
	guitars = append(guitars, "guitas 1")
	guitars = append(guitars, "guitas 2")
	// pemanggilann data
	fmt.Println(guitars[0])

	// deklarasi kedua
	colors := []string{"hitam", "putih", "biru"}
	// pemangilan data
	for _, color := range colors {
		fmt.Println(color)
		if color == "biru" {
			fmt.Println("warna kesukaan saya")
		}
	}

	// fungsi cap dan len pada slice
	// fungsi len akan mengukur panjang/jumlah data yang diambil
	// fungsi cap akan mengukur kapasitas dari data
	fmt.Println(len(colors))
	fmt.Println(cap(colors))

	fmt.Println("===============")
	fmt.Println(len(colors[0:2]))
	fmt.Println(cap(colors[0:2]))
	fmt.Println(colors[0:2])

	fmt.Println("===============")
	fmt.Println(len(colors[1:3]))
	fmt.Println(cap(colors[1:3]))
	fmt.Println(colors[1:3])

	// fungsi copy
	// fungsi copy memiliki 2 parameter
	// copy (dst, src)
	// dst => tujuan
	// src => asal

	// buat array tipe string dengan jumlah element yang bisa ditampung sebanyak 3 sebagai tujuan
	platNomor := make([]string, 4)
	// buat array sebagai asal
	platNomorBackup := []string{"AB 189 BC", "H 789 BNN", "R 345 ZS", "G 890 PO"}
	//copy akan mengembalikan angka yang mana menunjukan jumlah data yang tercopy
	copy := copy(platNomor, platNomorBackup)
	fmt.Println("==================")
	fmt.Println(copy)
	fmt.Println("len dari tujuan : ", len(platNomor))
	for _, plat := range platNomor {
		fmt.Println(plat)
	}
	fmt.Println("==================")

	// ++++++++++++++++ MAP  ++++++++++++++++
	// jika pada bahasa lain map ini sebagai array assosiative

	// deklarasi pertama
	// dengan syntax
	// var namaMap map[tipe index]tipe value
	var konten map[string]int
	// inisiasi
	konten = map[string]int{}
	// memasukan data ke map
	konten["PHP"] = 10
	konten["JAVA"] = 9
	konten["GO"] = 10

	fmt.Println(konten)
	fmt.Println(konten["go"])
	fmt.Println(konten["GO"])
	for key, bahasa := range konten {
		fmt.Println("key : ", key, "value : ", bahasa)
	}
	fmt.Println("=============")

	// dekralasi kedua
	books := map[string]string{
		"Novel": "Novel1",
		"Comic": "Comic1",
		"Sains": "Sains",
	}
	fmt.Println(books)
	fmt.Println(books["Sains"])
	for key, value := range books {
		fmt.Println("key : ", key, "value : ", value)
	}

	fmt.Println("================")
	// hapus data dari map
	delete(books, "Comic")
	for key, value := range books {
		fmt.Println("key : ", key, "value : ", value)
	}

	// cara cek apakah ada index tertentu di dalam map
	// sama seperti pada perulangan kita bisa menggukanan _ jiga ada parameter yang tidak ingin dipakai
	_, isAvailable := books["Comic"]
	fmt.Println(isAvailable)

	value, isAvailable := books["Sains"]
	fmt.Println(isAvailable)
	fmt.Println(value)

	// ++++++++++++++++ SLICE MAP  ++++++++++++++++
	myKey := []map[string]string{
		{"key1": "Rahasia", "unlock": "unlockRahasia"},
		{"key2": "Rahasia123", "unlock": "unlockRahasi123"},
		{"key3": "Rahasia12345", "unlock": "unlockRahasi12345"},
	}

	for _, item := range myKey {
		fmt.Println(item["key1"], item["unlock"])
	}

	// ++++++++++++++++ QUIZ  ++++++++++++++++
	// hitung rata-rata dari
	fmt.Println("===============")
	// nilai awal
	scores := [...]int{100, 80, 75, 92, 70, 93, 88, 67}
	// nilai hasil penjumlahan
	totals := 0
	for _, item := range scores {
		totals = totals + item
	}
	length := len(scores)
	avg := float64(totals) / float64(length)
	fmt.Println(avg)

	// lalu masukan nilai dari scores di atas ke dalam slice dengan syarat score > 90

	// solusi kedua
	var goodScores []int
	for _, item := range scores {
		if item > 90 {
			// solusi pertama
			// goodScores := []int{item}
			// fmt.Println(goodScores)

			// solusi kedua
			goodScores = append(goodScores, item)
		}
	}

	// solusi kedua
	fmt.Println(goodScores)

	// Pemanggian function biasa
	Sepeda()

	// Pemanggian function dengan parameter
	Human("Reza")

	// Pemanggian function dengan output
	// harus ditampung di varible sebagai tempat menyimpan outputnya
	roar := Macan("Bagaimana Suara Macan ???")
	fmt.Println(roar)

	// Pemanggilan function penjumlahan
	add := add(3, 5)
	fmt.Println(add)

	// Pemanggilan function pengurangan
	kurang := minus(10, 8)
	fmt.Println(kurang)

	// Pemanggilan function lingkaran dengan 2 return
	// semua return harus masuk (keliling dan luas)
	// jika ingin memilih salah satu maka pakai _ saja
	var r float64 = 5
	// var keliling, luas = lingkaran(r)
	var keliling, _ = lingkaran(r)
	fmt.Println("r = ", r, "Keliling lingaran", keliling)
	// fmt.Println("r = ", r, "Luas lingaran", luas)

	// Pemanggilan function persegi
	p := 1
	l := 10
	var kelilingP, luasP = persegi(p, l)
	fmt.Println("P = ", p, "L = ", l, "Keliling = ", kelilingP)
	fmt.Println("P = ", p, "L = ", l, "Luas = ", luasP)

	final, greet := grade(20, 10)
	fmt.Println(final, greet)

	pria, wanita := gender("reza")
	fmt.Println(pria)
	fmt.Println(wanita)

	arr := []int{2, 3, 10}
	totalArr := sumArray(arr)
	fmt.Println(totalArr)
	fmt.Println("=======")
	res, err := kalkulator(5, 3, ";")
	if err != nil {
		fmt.Println("Terjadi Kesalahan")
		fmt.Println(err)
	}
	fmt.Println(res)

	// ++++++++++++++++ STRUCT  ++++++++++++++++
	// ini pemanggilan struct dari deklarasi di atas
	fmt.Println("========= STRUCT =========")
	user1 := User{}
	// mengisi data ke properties yang ada di struct User
	user1.ID = 1
	user1.FirstName = "Prakadinata"
	user1.LastName = "Mahendra"
	user1.Email = "prada@gmail.com"
	user1.IsActive = true

	user2 := User{}
	// mengisi data ke properties yang ada di struct User
	// ini boleh acak pengsisian properties nya karena kita deklarasikan properties nya langsung
	user2.ID = 2
	user2.LastName = "Putri"
	user2.FirstName = "Selenia"
	user2.IsActive = true
	user2.Email = "selen@gmail.com"

	// ini boleh acak pengsisian properties nya karena kita deklarasikan properties nya langsung
	user3 := User{FirstName: "Reza", ID: 3, Email: "reza@gmail.com", LastName: "Wijaya", IsActive: true}

	// ini pengisian data nya harus urut karena properties tidak kita deklarasikan
	user4 := User{
		4,
		"Abdas",
		"Reza",
		"abdas@gmail.com",
		true,
	}

	// print semua data
	fmt.Println(user1)
	fmt.Println(user2)
	fmt.Println(user3)
	fmt.Println(user4)

	// print data tertentu
	fmt.Println("ID user pertama", user1.ID)
	fmt.Println("Email user kedua", user2.Email)
	fmt.Println("First Name user ketiga", user3.FirstName)
	fmt.Println("IsActive user keempat", user4.IsActive)

	// struct sebagai parameter
	fmt.Println("=========== Stuct sebagai parameter ===========")
	showUser1 := showUser(user1)
	showUser2 := showUser(user4)
	fmt.Println(showUser1)
	fmt.Println(showUser2)

	// ++++++++++++++++ EMBEDDED STRUCT  ++++++++++++
	// kita masukan isi dulu ke dalam struct mahasiswa
	reza := mahasiswa{
		Nama:     "Reza",
		Nim:      19102149,
		Fakultas: "Fakultas Informatika",
		Jurusan:  "Teknik Informatika",
		Kelas:    "S1IF-07-P",
		Usia:     21,
	}

	irfan := mahasiswa{
		Nama:     "Irfan",
		Nim:      19102148,
		Fakultas: "Fakultas Informatika",
		Jurusan:  "Rekayasa Perangkat Lunak",
		Kelas:    "S1IF-06-N",
		Usia:     23,
	}

	wijaya := mahasiswa{
		Nama:     "Wijaya",
		Nim:      19102147,
		Fakultas: "Fakultas Informatika",
		Jurusan:  "Teknik Informatika",
		Kelas:    "S1IF-08-H",
		Usia:     25,
	}

	// jika struct mahasiswa sudah ada isi nya lanjut mengisi data ke struct grup
	// namun sebelum itu kita harus mengisi dulu slice anggota group
	anggota := []mahasiswa{irfan, wijaya}
	kelasMM4 := group{
		Nama: "Kelas MM4 Aweeeu",
		// kita pilih siapa yang menjadi admin
		Admin: reza,
		// karena ini slice maka harus kita definiskan terlebih dahulu diatas
		Anggota:     anggota,
		isAvailable: true,
	}

	// kita cetak kelasMM4
	// fmt.Println(kelasMM4)
	// agar lebih rapih kita pakai function saja

	groupKelasMM4(kelasMM4)

	// ++++++++++++++++ METHODE STRUCT  ++++++++++++
	// isi data ke struct nation
	Indonesia := Nation{
		Nama:     "Indonesia",
		Penduduk: 12000000,
		Luas:     123,
		Presiden: "Jokowi",
		Benua:    "Asia",
	}

	Rusia := Nation{
		Nama:     "Rusia",
		Penduduk: 100000,
		Luas:     200,
		Presiden: "Vladimir Putin",
		Benua:    "Eropa",
	}

	Malaysia := Nation{
		Nama:     "Malaysia",
		Penduduk: 110000,
		Luas:     345,
		Presiden: "Mahathir Mohamad",
		Benua:    "Asia",
	}

	KoreaUtara := Nation{
		Nama:     "Korea Utara",
		Penduduk: 178000,
		Luas:     506,
		Presiden: "Kim Jong Un",
		Benua:    "Asia",
	}

	// function show saya deklarasikan diatas
	showID := Nation.show(Indonesia)
	showRSA := Nation.show(Rusia)
	showMLY := Nation.show(Malaysia)
	showKUT := Nation.show(KoreaUtara)

	fmt.Println(showID)
	fmt.Println(showRSA)
	fmt.Println(showMLY)
	fmt.Println(showKUT)

	// ++++++++++++++++ QUIZ STRUCT  ++++++++++++
	Eza := Siswa{
		Nama:     "Reza",
		Kelas:    "XII MIPA 1",
		Angkatan: 2016,
		Motivasi: "Ingin menambah wawasan",
	}

	Jhon := Siswa{
		Nama:     "Jhon",
		Kelas:    "XI IPS 5",
		Angkatan: 2017,
		Motivasi: "Ingin menambah koneksi",
	}

	Michelle := Siswa{
		Nama:     "Michelle",
		Kelas:    "XII MIPA 2",
		Angkatan: 2016,
		Motivasi: "Memperdalam public speaking",
	}

	Angeline := Siswa{
		Nama:     "Angeline",
		Kelas:    "XI MIPA 3",
		Angkatan: 2017,
		Motivasi: "Ingin menghilangkan rasa grogi",
	}

	Abdul := Siswa{
		Nama:     "Abdul",
		Kelas:    "XI IPS 1",
		Angkatan: 2017,
		Motivasi: "Memenuhi janji kepada diri sendiri",
	}

	Mayer := Siswa{
		Nama:     "Mayer",
		Kelas:    "XII IPS 3",
		Angkatan: 2016,
		Motivasi: "Mengasah kemampuan kepemimpinan",
	}

	anggotaOsis := []Siswa{Eza, Angeline, Michelle}

	GroupOsis := Osis{
		Nama:      "SMA NEGERI 1 WAKANDA",
		Ketua:     Jhon.Nama,
		Wakil:     Mayer.Nama,
		Bendahara: Abdul.Nama,
		Anggota:   anggotaOsis,
	}

	Osis.showOsis(GroupOsis)

	// imported struc from another package
	buku1 := buku.Buku{
		Judul:     "Buku Satu",
		Halaman:   349,
		Pengarang: "Mr.A",
		Penerbit:  "Erlangga",
	}
	fmt.Println(buku1)

}

// ++++++++++++++++ FUNCTION  ++++++++++++++++
// function biasa
func Sepeda() {
	fmt.Println("saya dari function sepeda")
}

// function dengan parameter
func Human(nama string) {
	fmt.Print(nama)
}

// function dengan output
func Macan(voice string) string {
	ask := voice + "ROARRRRRR !!!!"
	return ask
}

// function penjumlahan
func add(num1 int, num2 int) int {
	// res := num1 + num2
	// return res

	return num1 + num2
}

// function dengan tipe data parameter yang sama
func minus(num1, num2 int) int {
	return num1 - num2
}

// function dengan multiple return
func lingkaran(r float64) (float64, float64) {
	// function ini akan menghitung luas dan keliling lingkaran
	keliling := (22 / 7) * 2 * r
	luas := (22 / 7) * r * r

	return keliling, luas
}

// function dengan multiple return
func persegi(panjang, lebar int) (int, int) {
	keliling := 2 * (panjang + lebar)
	luas := panjang * lebar

	return keliling, luas

}

func grade(absensi, kuis int) (int, string) {
	final := absensi + kuis
	if final > 70 {
		return final, fmt.Sprintln("Kamu Lulus !!!")
	} else {
		return final, fmt.Sprintln("Anda belum lulus !!!")
	}
}

// predifined return values
// menyebut nama varibale pada nilai balikan
func gender(nama string) (pria string, wanita string) {
	if nama == "reza" {
		pria = "Anda Pria"

	} else {
		wanita = "Anda Wanita"

	}

	return
}

// ++++++++++++++++ QUIZ FUNCTION  ++++++++++++++++
func sumArray(arr []int) int {
	num := 0
	for _, item := range arr {
		num += item
	}
	return num
}

func kalkulator(num1, num2 float64, operator string) (hasil float64, err error) {
	switch operator {
	case "+":
		hasil = num1 + num2

	case "-":
		hasil = num1 - num2

	case "*":
		hasil = num1 * num2

	case "/":
		hasil = num1 / num2

	default:
		err = errors.New("Unknown Operation")
	}
	return
}

// struct sebagai parameter
// jadi kita set struct user menjadi tipe data dari parameter
func showUser(user User) (res string) {
	// nah user ini sudah merepresentasikan struct User jadi kita bisa mengambil properties nya
	// %s akan mewakili string
	res = fmt.Sprintf("Nama: %s %s , Email %s", user.FirstName, user.LastName, user.Email)
	// kita mengharapkan output = Nama : Putri Selenia Email : putri @gmail.com
	return
}

// struct embedded show data function
func groupKelasMM4(groupWA group) {

	fmt.Printf("\n\nNama Group : %s\n", groupWA.Nama)
	fmt.Printf("Nama Admin : %s\n", groupWA.Admin.Nama)
	fmt.Printf("Jumlah Anggota : %d", len(groupWA.Anggota))

	fmt.Println("\nList Anggota : ")

	for _, item := range groupWA.Anggota {
		fmt.Printf("\nNama : %s\n", item.Nama)
		fmt.Printf("Jurusan : %s\n", item.Jurusan)
		fmt.Printf("Fakultas : %s\n", item.Fakultas)
		fmt.Printf("Kelas : %s\n", item.Kelas)
		fmt.Printf("Nim : %d\n", item.Nim)
		fmt.Printf("Usia : %d\n\n", item.Usia)
	}
}
