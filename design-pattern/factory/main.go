package main

func main() {

	// kita membuat object dari interface content creator
	var creator ContentCreator = &RezaIrfan{}

	konten := creator.CreateContent()
	konten.Play()
}
