package product

// ini adalah kursi merk vienna dari informa
type vienna struct{}

func (v vienna) Price() float64 {
	return 1449000
}

func (v vienna) IsIotEnable() bool {
	return false
}

func (v vienna) IsSoft() bool {
	return true
}
