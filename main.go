package main

import "fmt"

const englishHelloPrefix = "Hello, "

const spanishHelloPrefix = "Hola, "

func Hello(name string, lang string) string {

	if name == "" {
		name = "World"
	}

	prefix := englishHelloPrefix

	switch lang {
	case "Spanish":
		prefix = spanishHelloPrefix
	case "English":
		prefix = englishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return prefix + name
}

func main() {
	fmt.Println(Hello("Marko", "English"))

}
