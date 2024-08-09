package main

import "fmt"

func main() {
	fmt.Println(Hello("world", ""))
}

const englishPrefix = "Hello, "
const spanishPrefix = "Hola, "
const frenchPrefix = "Bonjour, "
const spanishSufix = "Mundo"
const frenchSufix = "Monde"
const englishSufix = "World"

func Hello(name, language string) string {

	if language == "French" {
		if name == "" {
			name = frenchSufix
		}
		return frenchPrefix + name + "!"
	}

	if language == "Spanish" {
		if name == "" {
			name = spanishSufix
		}
		return spanishPrefix + name + "!"
	}

	if name == "" {
		name = englishSufix
	}

	return englishPrefix + name + "!"
}
