package main

import "testing"

// nama function harus sama dengan nama function yang akan di test dan harus diawali Test

// ini unit test tanpa param di function Hello
func TestHello(t *testing.T) {
	// function yang akan kita uji
	yangDidapat := Hello()

	// output yang diharapkan
	yangDiinginkan := "Hello Unit Test"

	// cek apakah output yang didapat sama dengan output yang diinginkan
	if yangDidapat != yangDiinginkan {
		t.Errorf("yang didapat %s, yang diinginkan %s", yangDidapat, yangDiinginkan)
	}
}

// cara run nya: go test
// atau go test hello.go hello_test.go

// ini unit test dengan param di function HelloWithName
func TestHelloWithName(t *testing.T) {
	// karena kita ada 2 test (kondisi, ada if else), maka kita buat 2 sub test

	// sub pertama (jika name kosong)
	// syntax: t.Run("nama_sub_test", func(t *testing.T))
	t.Run("Say hello to friend", func(t *testing.T) {
		yangDidapat := HelloWithName("")

		yangDiinginkan := "Hello Friend"

		if yangDidapat != yangDiinginkan {
			t.Errorf("yang didapat %s, yang diinginkan %s", yangDidapat, yangDiinginkan)
		}
	})

	// sub kedua (jika name tidak kosong)
	// syntax: t.Run("nama_sub_test", func(t *testing.T))
	t.Run("Say hello to someone", func(t *testing.T) {
		yangDidapat := HelloWithName("Reza")

		yangDiinginkan := "Hello Reza"

		if yangDidapat != yangDiinginkan {
			t.Errorf("yang didapat %s, yang diinginkan %s", yangDidapat, yangDiinginkan)
		}
	})

}
