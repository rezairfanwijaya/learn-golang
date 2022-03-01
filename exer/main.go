package main

import "fmt"

func main() {
	// bikin varible
	var nama string = "Nama Anda"
	fmt.Println(nama)

	var usia int
	usia = 80
	fmt.Println(usia + 90)

	alamat := "Indonesia"
	fmt.Println(alamat)

	alamat = "jawa"
	fmt.Println(alamat)

	if usia > 30 {
		fmt.Println("usia anda", usia, "anda tua")
	}

	if usia > 30 && usia < 50 {
		fmt.Println("usia anda paruh baya")

		if usia > 35 && usia < 41 {
			fmt.Println("nice")
		}
	}

	switch usia {
	case 30:
		fmt.Println("anda 30")
	case 40:
		fmt.Println("anda 40")
	default:
		fmt.Println("diluar jangkuan")
	}

	iteration := 50
	for i := 1; i <= iteration; i++ {
		if i%2 == 0 {
			fmt.Println(i, "= genap")
		} else {
			fmt.Println(i)
		}
	}

	var negara = "amerika"
	for index, huruf := range negara {
		fmt.Println(index, "=", string(huruf))
	}

	size := 1
	for size <= 35 {
		fmt.Println(size)
		size++
	}

	var n [3]string
	n[0] = "Indonesia"
	n[1] = "Amerika"
	n[2] = "Rusia"
	fmt.Println(n)
	for _, item := range n {
		fmt.Println(item)
	}

	var key = [2]int{3, 5}
	fmt.Println(key)

	age := [6]int{1, 4, 2, 6, 7, 2}
	for _, item := range age {
		fmt.Println(item)
	}

	bahasa := [...]string{
		"indo",
		"jawa",
		"sunda",
		"badui",
	}

	fmt.Println(len(bahasa))
	fmt.Println(cap(bahasa))
	for _, item := range bahasa {
		fmt.Println(item)
	}

	mouse := make([]string, 5)
	mouse[0] = "Logitech1"
	mouse[1] = "Logitech2"
	mouse[2] = "Logitech3"
	fmt.Println(mouse)
	fmt.Println(len(mouse))
	m := mouse[0:2]
	fmt.Println(m)
	fmt.Println(len(m))
	fmt.Println(cap(m))

	var piano []string
	piano = append(piano, "piano1")
	piano = append(piano, "piano2")
	piano = append(piano, "piano")
	for index, items := range piano {
		fmt.Println(index, items)
	}

	colors := []string{"red", "green", "blue"}
	fmt.Println(colors)
	newColor := colors[0:2]
	fmt.Println(newColor)
	fmt.Println(len(newColor))
	fmt.Println(cap(newColor))
	fmt.Println("===============")
	hairColor := colors[1:3]
	fmt.Println(hairColor)
	fmt.Println(len(hairColor))
	fmt.Println(cap(hairColor))

	container := make([]string, 4)
	actors := []string{"Actor1", "Actor2", "Actor3", "Actor4", "Actor5"}
	fmt.Println(actors)
	nobel := copy(container, actors)
	fmt.Println(nobel)
	fmt.Println(container)
	fmt.Println("=========")

	newContainer := container[0:2]
	fmt.Println(newContainer)
	fmt.Println(len(newContainer))
	fmt.Println(cap(newContainer))
	fmt.Println("=========")

	oldContainer := container[1:3]
	fmt.Println(oldContainer)
	fmt.Println(len(oldContainer))
	fmt.Println(cap(oldContainer))

}