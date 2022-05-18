package product

import (
	abstract_factory "myapp/design-pattern/abstract-factory/abstract"
)

// ini adalah meja merk magnum dari informa
type magnum struct{}

func (m magnum) Price() float64 {
	return 1190000
}

func (m magnum) Size() abstract_factory.Dimension {
	return abstract_factory.Dimension{Length: 100, Width: 100, Height: 100}
}

func (m magnum) IsFoldable() bool {
	return false
}
