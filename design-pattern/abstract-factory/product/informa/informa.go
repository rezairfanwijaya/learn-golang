package product

import (
	abstract_factory "myapp/design-pattern/abstract-factory/abstract"
)

type Informa struct{}

// lalu kita implementasikan factory nya
func (i Informa) CreateChair() abstract_factory.Chair {
	return vienna{}
}

func (i Informa) CreateTable() abstract_factory.Table {
	return magnum{}
}

func (i Informa) CreateCupboard() abstract_factory.Cupboard {
	return verona{}
}
