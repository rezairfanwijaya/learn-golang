package product

import (
	abstract_factory "myapp/design-pattern/abstract-factory/abstract"
)

// ini adalah lemari merk verona dari informa
type verona struct{}

func (v verona) Price() float64 {
	return 12599000
}

func (v verona) Size() abstract_factory.Dimension {
	return abstract_factory.Dimension{Length: 200, Width: 200, Height: 200}
}

func (v verona) IsIotEnable() bool {
	return false
}
