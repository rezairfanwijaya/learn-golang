package product

import (
	abstract_factory "myapp/design-pattern/abstract-factory/abstract"
)

type Ikea struct{}

// lalu kita implementasikan factory nya
func (i Ikea) CreateChair() abstract_factory.Chair {
	return adde{}
}

func (i Ikea) CreateTable() abstract_factory.Table {
	return lisabbo{}
}

func (i Ikea) CreateCupboard() abstract_factory.Cupboard {
	return brimnes{}
}
