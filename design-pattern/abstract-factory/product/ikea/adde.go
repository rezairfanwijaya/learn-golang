package product

// ini adalah kursi merk adde dari ikea
type adde struct{}

func (a adde) Price() float64 {
	return 180000
}

func (a adde) IsIotEnable() bool {
	return false
}

func (a adde) IsSoft() bool {
	return false
}
