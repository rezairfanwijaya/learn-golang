package product

import (
	abstract_factory "myapp/design-pattern/abstract-factory/abstract"
)

// ini adalah meja merk lisabbo dari ikea
type lisabbo struct{}

func (l lisabbo) Price() float64 {
	return 2900000
}

func (l lisabbo) Size() abstract_factory.Dimension {
	return abstract_factory.Dimension{Length: 105, Width: 105, Height: 74}
}

func (l lisabbo) IsFoldable() bool {
	return false
}
