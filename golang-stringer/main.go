package main

import "fmt"

// i/o stringer adalah mengcustom metho println itu sendiri
// sama seperti custom error

type cat struct {
	name  string
	sound string
}

type dog struct {
	name  string
	sound string
}

func main() {
	dog := dog{
		name:  "bobol",
		sound: "bobo",
	}

	cat := cat{
		name:  "letto",
		sound: "meow",
	}

	fmt.Printf("\nTanpa i/o stringer\n")
	fmt.Println(dog)

	fmt.Printf("\nDengan i/o stringer\n")
	fmt.Println(cat)
}

// function untuk melakukan custom i/o stringer
func (c cat) String() string {
	return fmt.Sprintf("%s say %s", c.name, c.sound)
}
