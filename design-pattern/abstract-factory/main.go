package main

import (
	"fmt"
	product_ikea "myapp/design-pattern/abstract-factory/product/ikea"
	product_informa "myapp/design-pattern/abstract-factory/product/informa"
)

func main() {
	ikea := product_ikea.Ikea{}
	chair_ikea := ikea.CreateChair()
	fmt.Println(chair_ikea.IsIotEnable(), chair_ikea.IsSoft(), chair_ikea.Price())

	informa := product_informa.Informa{}
	chair_informa := informa.CreateChair()
	fmt.Println(chair_informa.IsIotEnable(), chair_informa.IsSoft(), chair_informa.Price())

}
