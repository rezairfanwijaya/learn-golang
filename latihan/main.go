package main

import (
	"fmt"
)

func main() {
	// menulis varibale
	var nama string = "Nama saya"
	fmt.Println(nama)

	var angka int
	angka = 80
	fmt.Println("Ada angka : ", angka)

	kalimat := "Hallo saya angka"
	fmt.Println(kalimat, angka)

	// percabangan if
	if angka > 50 {
		fmt.Println("anda lulus")
	}

	if angka > 80 && angka < 100 {
		fmt.Println("nilai pas")
	}

	// switch case
	switch angka {
	case 90:
		fmt.Println("aduh")
	default:
		fmt.Println("harus 90")
	}

	// gabungan switch case dan if
	if angka < 70 {
		fmt.Println("belum lulus")
	} else if angka >= 80 {
		fmt.Println("lulus")
		switch angka {
		case 80:
			{
				fmt.Println("dengan nilai A")
			}
		case 90:
			{
				fmt.Println("dengan nilai A+")
			}
		}
		
	}

	// looping
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	for index, sentences := range kalimat {
		fmt.Println("huruf ke", index+1, ":", string(sentences))
	}

	t := 1
	for t < 10 {
		fmt.Println("While ke-", t)
		t++
	}

	// array
	var rumah [3]string
	rumah[0] = "Rumah Pertama"
	rumah[1] = "Rumah Kedua"
	rumah[2] = "Rumah Ketiga"
	fmt.Println(rumah)
	fmt.Println(len(rumah))
	fmt.Println(rumah[2])

	kursi := [3]string{
		"kuris pertama",
		"kuris kedua",
		"kuris ketiga",
	}

	fmt.Println(kursi)
	fmt.Println(kursi[2])
	fmt.Println(len(kursi))

	meja := [...]string{
		"Meja 1",
		"Meja 2",
		"Meja 3",
		"Meja 4",
		"Meja 5",
		"Meja 6",
	}

	fmt.Println(meja)
	fmt.Println(len(meja))

}
