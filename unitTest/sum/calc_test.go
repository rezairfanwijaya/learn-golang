package main

import (
	"testing"
)

func TestTambah(t *testing.T) {
	// yang didapat
	yangDidapat := tambah(2, 2)

	// yang diinginkan
	yangDiinginkan := 4

	if yangDidapat != yangDiinginkan {
		// fmt.Println("Penjumlahan gagal")
		// t.Fail()

		t.Error("Penjumlahan gagal")
	}
}

func TestKurang(t *testing.T) {
	// yang didapat
	yangDidapat := kurang(2, 2)

	// yang diinginkan
	yangDiinginkan := 0

	if yangDidapat != yangDiinginkan {
		// fmt.Println("Pengurangan gagal")
		// t.Fail()

		t.Error("Pengurangan gagal")
	}
}

func TestKali(t *testing.T) {
	// yang didapat
	yangDidapat := kali(2, 2)

	// yang diinginkan
	yangDiinginkan := 4

	if yangDidapat != yangDiinginkan {
		// fmt.Println("Perkalian gagal")
		// t.Fail()

		t.Error("Perkalian gagal")
	}
}

func TestBagi(t *testing.T) {
	// yang didapat
	yangDidapat := bagi(2, 2)

	// yang diinginkan
	yangDiinginkan := 1

	if yangDidapat != yangDiinginkan {
		// fmt.Println("Pembagian gagal")
		// t.Fail()

		t.Error("Pembagian gagal")
	}
}
