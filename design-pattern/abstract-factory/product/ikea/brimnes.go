package product

import (
	abstract_factory "myapp/design-pattern/abstract-factory/abstract"
)

// ini adalah lemari merk brimnes dari ikea
type brimnes struct{}

func (b brimnes) Price() float64 {
	return 2799000
}

func (b brimnes) Size() abstract_factory.Dimension {
	return abstract_factory.Dimension{Length: 200, Width: 200, Height: 200}
}

func (b brimnes) IsIotEnable() bool {
	return false
}
