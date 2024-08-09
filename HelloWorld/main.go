package main

import (
	"strings"
)

func main() {

}

const (
	englishPrefix = "Hello, "
	spanishPrefix = "Hola, "
	frenchPrefix  = "Bonjour, "
	spanishSufix  = "Mundo"
	frenchSufix   = "Monde"
	englishSufix  = "World"
)

const (
	english = "english"
	french  = "french"
	spanish = "spanish"
)

func Hello(name, language string) string {

	return greetingPrefix(language) + managerName(language, name) + "!"
}

func greetingPrefix(language string) string {
	var prefix = ""
	language = strings.ToLower(language)

	switch language {
	case english:
		prefix = englishPrefix
	case french:
		prefix = frenchPrefix
	case spanish:
		prefix = spanishPrefix
	default:
		prefix = englishPrefix
	}
	return prefix
}

func managerName(language, name string) string {
	if name != "" {
		return name
	}

	var sufix = ""
	language = strings.ToLower(language)

	switch language {
	case english:
		sufix = englishSufix
	case french:
		sufix = frenchSufix
	case spanish:
		sufix = spanishSufix
	default:
		sufix = englishSufix
	}
	return sufix
}
