package variables

import "fmt"

func Kata() string {
	// syntax
	// var namaVariable tipedata = "value"
	// var nama string = "Reza Irfan Wijaya"

	// atau bisa juga dengan menulis sperti ini
	nama := "Reza Irfan Wijaya"
	// cara merubah nilai yg di definisikan oleh :=
	nama = "Reza irfan"
	return fmt.Sprintf(nama)

}
