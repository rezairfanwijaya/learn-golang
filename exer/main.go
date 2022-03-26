package main

import (
	"errors"
	"fmt"
)

func main() {
	// bikin varible
	var nama string = "Nama Anda"
	fmt.Println(nama)

	var usia int
	usia = 80
	fmt.Println(usia + 90)

	alamat := "Indonesia"
	fmt.Println(alamat)

	alamat = "jawa"
	fmt.Println(alamat)

	if usia > 30 {
		fmt.Println("usia anda", usia, "anda tua")
	}

	if usia > 30 && usia < 50 {
		fmt.Println("usia anda paruh baya")

		if usia > 35 && usia < 41 {
			fmt.Println("nice")
		}
	}

	switch usia {
	case 30:
		fmt.Println("anda 30")
	case 40:
		fmt.Println("anda 40")
	default:
		fmt.Println("diluar jangkuan")
	}

	iteration := 50
	for i := 1; i <= iteration; i++ {
		if i%2 == 0 {
			fmt.Println(i, "= genap")
		} else {
			fmt.Println(i)
		}
	}

	var negara = "amerika"
	for index, huruf := range negara {
		fmt.Println(index, "=", string(huruf))
	}

	size := 1
	for size <= 35 {
		fmt.Println(size)
		size++
	}

	var n [3]string
	n[0] = "Indonesia"
	n[1] = "Amerika"
	n[2] = "Rusia"
	fmt.Println(n)
	for _, item := range n {
		fmt.Println(item)
	}

	var key = [2]int{3, 5}
	fmt.Println(key)

	age := [6]int{1, 4, 2, 6, 7, 2}
	for _, item := range age {
		fmt.Println(item)
	}

	bahasa := [...]string{
		"indo",
		"jawa",
		"sunda",
		"badui",
	}

	fmt.Println(len(bahasa))
	fmt.Println(cap(bahasa))
	for _, item := range bahasa {
		fmt.Println(item)
	}

	mouse := make([]string, 5)
	mouse[0] = "Logitech1"
	mouse[1] = "Logitech2"
	mouse[2] = "Logitech3"
	fmt.Println(mouse)
	fmt.Println(len(mouse))
	m := mouse[0:2]
	fmt.Println(m)
	fmt.Println(len(m))
	fmt.Println(cap(m))

	// slice

	var piano []string
	piano = append(piano, "piano1")
	piano = append(piano, "piano2")
	piano = append(piano, "piano")
	for index, items := range piano {
		fmt.Println(index, items)
	}

	colors := []string{"red", "green", "blue"}
	fmt.Println(colors)
	newColor := colors[0:2]
	fmt.Println(newColor)
	fmt.Println(len(newColor))
	fmt.Println(cap(newColor))
	fmt.Println("===============")
	hairColor := colors[1:3]
	fmt.Println(hairColor)
	fmt.Println(len(hairColor))
	fmt.Println(cap(hairColor))

	// copy

	container := make([]string, 4)
	actors := []string{"Actor1", "Actor2", "Actor3", "Actor4", "Actor5"}
	fmt.Println(actors)
	nobel := copy(container, actors)
	fmt.Println(nobel)
	fmt.Println(container)
	fmt.Println("=========")

	newContainer := container[0:2]
	fmt.Println(newContainer)
	fmt.Println(len(newContainer))
	fmt.Println(cap(newContainer))
	fmt.Println("=========")

	oldContainer := container[1:3]
	fmt.Println(oldContainer)
	fmt.Println(len(oldContainer))
	fmt.Println(cap(oldContainer))

	// map
	var myMap map[string]int
	myMap = map[string]int{}
	myMap["a"] = 1
	myMap["b"] = 2
	myMap["c"] = 3
	fmt.Println(myMap)
	for index, item := range myMap {
		fmt.Println("index : ", index, "value : ", item)
	}

	Nation := map[string]string{
		"Asia":      "Indonesia",
		"Eropa":     "Inggris",
		"Australia": "Australia",
	}

	for _, item := range Nation {
		fmt.Println(item)
	}

	delete(Nation, "Eropa")
	fmt.Println("====================")
	for _, item := range Nation {
		fmt.Println(item)
	}

	item, isAvailable := Nation["Eropa"]
	fmt.Println(item)
	fmt.Println(isAvailable)

	// quiz
	fmt.Println("=========== QUIZ ===========")
	Scores := [3]int{10, 9, 1}
	var tmp int
	for _, item := range Scores {
		tmp += item
	}

	result := float64(tmp) / float64(len(Scores))
	fmt.Println(result)

	// function
	human()

	tebakSuara("ROARRRRR !!!")

	add := add(3, 10)
	fmt.Println(add)

	say := greeting("reza")
	fmt.Println(say)

	luas, keliling := persegi(10, 10)
	fmt.Println("luas persegi : ", luas)
	fmt.Println("keliling persegi : ", keliling)

	// quiz
	fmt.Println("============ QUIZ ============")
	grade := []int{10, 0, 10}
	res := sumArr(grade)
	fmt.Println(res)

	result, err := calc(10, 3, "/")
	if err != nil {
		fmt.Println("Terjadi kesalahan")
		fmt.Println(err)
	}
	fmt.Println(result)

	human1 := humans{}
	human1.name = "Reza Irfan Wijaya"
	human1.address = "Dayeuhluhur"
	human1.age = 22

	human2 := humans{}
	human2.name = "Reza Irfan Wijaya"
	human2.address = "Dayeuhluhur"
	human2.age = 22

	human3 := humans{name: "Abdas", age: 21, address: "Cilacap"}

	human4 := humans{
		name:    "Irfan",
		age:     19,
		address: "hanum",
	}

	human5 := humans{
		"Reza",
		23,
		"Sudimampir",
	}

	fmt.Println(human1)
	fmt.Println(human2)
	fmt.Println(human3)
	fmt.Println(human4)
	fmt.Println(human5)

	// struct sebagai parameter
	mhs1 := mahasiswa{
		nama:  "Reza Irfan",
		nim:   19102149,
		kelas: "S1-IF-07-P",
	}

	mhs2 := mahasiswa{}
	mhs2.nama = "Ipang"
	mhs2.kelas = "S1IF-03-T"
	mhs2.nim = 12345

	showMhs1 := displayMahasiswa(mhs1)

	showMhs2 := displayMahasiswa(mhs2)
	fmt.Println(showMhs1)
	fmt.Println(showMhs2)

	memberOne := Mahasiswa{
		Nama:     "Member One",
		Nim:      19102149,
		Kelas:    "S1IF-07-P",
		Fakultas: "Informatika",
		Jurusan:  "Teknik Informatika",
	}

	memberTwo := Mahasiswa{
		Nama:     "Member Two",
		Nim:      19102148,
		Kelas:    "S1IF-07-K",
		Fakultas: "Rekayasa Perangkat Lunak",
		Jurusan:  "Rekayasa Perangkat Lunak",
	}

	memberThree := Mahasiswa{
		Nama:     "Member Three",
		Nim:      19102147,
		Kelas:    "S1IF-07-D",
		Fakultas: "Informatika",
		Jurusan:  "Teknik Informatika",
	}

	memberFour := Mahasiswa{
		Nama:     "Member Four",
		Nim:      19102146,
		Kelas:    "S1IF-07-R",
		Fakultas: "Informatika",
		Jurusan:  "Teknik Informatika",
	}

	anggotaGroupMHS := []Mahasiswa{memberOne, memberTwo, memberFour}

	groupMHS := Group{
		Nama:    "Mahasiswa Kreatif Indonesia",
		Admin:   memberThree,
		Anggota: anggotaGroupMHS,
		IsAvail: true,
	}

	Group.showMember(groupMHS)

	// POINTER
	angka := 4
	alamatAngka := &angka

	fmt.Printf("Nilai dari angka : %d\n", angka)
	fmt.Println("Alamat nya : ", alamatAngka)

	*alamatAngka = 40

	fmt.Printf("Nilai dari angka : %d\n", angka)
	fmt.Println("Alamat nya : ", alamatAngka)

	num := 100
	fmt.Println("\n\nNilai Awal : ", num)
	fmt.Println("Memori Awal : ", &num)
	fmt.Println()
	exchange(&num, 900)

	var alumniReza Alumniku
	alumniReza.Nama = "Reza Irfan"
	alumniReza.Nim = 19102149

	alumniReza.change()
	fmt.Println(alumniReza.Nama)

}

// FUNCTION
// 1. tanpa parameter
func human() {
	fmt.Println("saya dipanggil dari function human")
}

// 2. dengan parameter
func tebakSuara(suara string) {
	fmt.Println("Suara yang dihasilakn ", suara)
}

// 3. dengan return
func add(a, b int) int {
	return a + b
}

func greeting(name string) string {
	return fmt.Sprintln("Hallo ", name)
}

// 4. dengan multiple return values
func persegi(l, p int) (int, int) {
	luas := l * p
	keliling := 2 * (p + l)
	return luas, keliling
}

// quiz function
func sumArr(arr []int) (total int) {
	total = 0
	for _, item := range arr {
		total += item
	}
	return
}

func calc(num1, num2 float64, operand string) (total float64, err error) {
	switch operand {
	case "+":
		total = num1 + num2
	case "-":
		total = num1 - num2
	case "*":
		total = num1 * num2
	case "/":
		total = num1 / num2
	default:
		err = errors.New("Operand Not Found")
	}

	return
}

// Struct
type humans struct {
	name    string
	age     int
	address string
}

// struct sebagai parameter
type mahasiswa struct {
	nama  string
	nim   int
	kelas string
}

func displayMahasiswa(mhs mahasiswa) string {
	return fmt.Sprintf("Nama : %s,  Nim : %d, Kelas : %s", mhs.nama, mhs.nim, mhs.kelas)
}

// embedded struct x quiz struct
type Mahasiswa struct {
	Nama     string
	Nim      int
	Kelas    string
	Fakultas string
	Jurusan  string
}

type Group struct {
	Nama    string
	Admin   Mahasiswa
	Anggota []Mahasiswa
	IsAvail bool
}

func (group Group) showMember() {
	fmt.Printf("\n\nSelamat Datang Di Grup %s\n", group.Nama)
	fmt.Printf("Admin Group : %s\n\n", group.Admin.Nama)
	fmt.Println("Anggota Group: ")

	for _, item := range group.Anggota {
		fmt.Printf("1. Nama : %s\n", item.Nama)
		fmt.Printf("2. Kelas : %s\n", item.Kelas)
		fmt.Printf("3. NIM : %d\n", item.Nim)
		fmt.Printf("4. Fakultas : %s\n", item.Fakultas)
		fmt.Printf("5. Jurusan : %s\n\n", item.Jurusan)
	}
}

// pointer sebagai parameter
func exchange(num *int, res int) {
	*num = res
	fmt.Println("Nilai Akhir : ", *num)
	fmt.Println("Memori Akhir: ", num)

}

type Alumniku struct {
	Nama string
	Nim  int
}

func (A *Alumniku) change() {
	A.Nama = A.Nama + " S.T"
}
