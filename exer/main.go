package main

import (
	"errors"
	"fmt"
)

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

	// slice

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

	// copy

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

	// map
	var myMap map[string]int
	myMap = map[string]int{}
	myMap["a"] = 1
	myMap["b"] = 2
	myMap["c"] = 3
	fmt.Println(myMap)
	for index, item := range myMap {
		fmt.Println("index : ", index, "value : ", item)
	}

	Nation := map[string]string{
		"Asia":      "Indonesia",
		"Eropa":     "Inggris",
		"Australia": "Australia",
	}

	for _, item := range Nation {
		fmt.Println(item)
	}

	delete(Nation, "Eropa")
	fmt.Println("====================")
	for _, item := range Nation {
		fmt.Println(item)
	}

	item, isAvailable := Nation["Eropa"]
	fmt.Println(item)
	fmt.Println(isAvailable)

	// quiz
	fmt.Println("=========== QUIZ ===========")
	Scores := [3]int{10, 9, 1}
	var tmp int
	for _, item := range Scores {
		tmp += item
	}

	result := float64(tmp) / float64(len(Scores))
	fmt.Println(result)

	// function
	human()

	tebakSuara("ROARRRRR !!!")

	add := add(3, 10)
	fmt.Println(add)

	say := greeting("reza")
	fmt.Println(say)

	luas, keliling := persegi(10, 10)
	fmt.Println("luas persegi : ", luas)
	fmt.Println("keliling persegi : ", keliling)

	// quiz
	fmt.Println("============ QUIZ ============")
	grade := []int{10, 0, 10}
	res := sumArr(grade)
	fmt.Println(res)

	result, err := calc(10, 3, "/")
	if err != nil {
		fmt.Println("Terjadi kesalahan")
		fmt.Println(err)
	}
	fmt.Println(result)

}

// FUNCTION
// 1. tanpa parameter
func human() {
	fmt.Println("saya dipanggil dari function human")
}

// 2. dengan parameter
func tebakSuara(suara string) {
	fmt.Println("Suara yang dihasilakn ", suara)
}

// 3. dengan return
func add(a, b int) int {
	return a + b
}

func greeting(name string) string {
	return fmt.Sprintln("Hallo ", name)
}

// 4. dengan multiple return values
func persegi(l, p int) (int, int) {
	luas := l * p
	keliling := 2 * (p + l)
	return luas, keliling
}

// quiz function
func sumArr(arr []int) (total int) {
	total = 0
	for _, item := range arr {
		total += item
	}
	return
}

func calc(num1, num2 float64, operand string) (total float64, err error) {
	switch operand {
	case "+":
		total = num1 + num2
	case "-":
		total = num1 - num2
	case "*":
		total = num1 * num2
	case "/":
		total = num1 / num2
	default:
		err = errors.New("Operand Not Found")
	}

	return
}
