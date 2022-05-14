package main

import (
	"myapp/design-pattern/factory"
)

func main() {

	// kita membuat object dari interface content creator
	var creator factory.ContentCreator
	creator = &factory.Eko{}

	konten := creator.CreateContent()
	konten.Play()
}
