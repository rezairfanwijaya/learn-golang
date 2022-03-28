package main

// tanpa params
func Hello() string {
	return "Hello Unit Test"
}

// with params
func HelloWithName(name string) string {
	if name == "" {
		return "Hello Friend"
	} else {
		return "Hello " + name
	}
}

func main() {
	println(Hello())
	println(HelloWithName("Reza"))
}
