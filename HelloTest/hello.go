package main

import "fmt"

func main() {
	fmt.Println(Hello("", ""))
}

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Holla, "

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	prefix := englishHelloPrefix
	if language == "Spanish" {
		prefix = spanishHelloPrefix
	}

	return prefix + name
}
