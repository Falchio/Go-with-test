package main

import (
	"fmt"
)

func main() {
	fmt.Println(Hello("", ""))
}

const (
	french  = "French"
	spanish = "Spanish"

	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Holla, "
	frenchHelloPrefix  = "Bonjour, "
)

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	prefix := greetingPrefix(language)
	return prefix + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix =frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}
