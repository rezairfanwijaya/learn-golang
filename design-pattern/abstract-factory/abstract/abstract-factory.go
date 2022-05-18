package abstract_factory

/*
	Abstarct factory ini akan membuat factory pattern
	pada case kali ini kita akan coba membuat sebuah factory pattern toko furniture(abstract factory) yang akan membuat factory (ikea,unifarm) yang akan menghasilkan produk kursi,meja dan lemari
	LETSS CODEEE !!
*/

// kita bikin produknya dulu (kursi,meja,lemari) --- interface
type Chair interface {
	Price() float64
	IsIotEnable() bool
	IsSoft() bool
}

type Dimension struct {
	Length, Width, Height int
}

type Table interface {
	Price() float64
	Size() Dimension
	IsFoldable() bool
}

type Cupboard interface {
	Price() float64
	Size() Dimension
	IsIotEnable() bool
}

// lalu kita bikin abstract factory nya --- interface
type FurnitureFactory interface {
	CreateChair() Chair
	CreateTable() Table
	CreateCupboard() Cupboard
}

// kemudian kita bikin factory nya --- struct

// factory akan mengimplentasikan furniture factory interface
// untuk mendapatakan produk kursi,meja,lemari kita harus bikin lagi struct kusi,meja,lemari yang mengimplemntasikan inteface chair,table,cupboard
