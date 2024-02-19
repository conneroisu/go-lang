package main

import "fmt"

func Hello(name, language string) string {
	if name == "" {
		name = "world"
	}
	prefix := greetingPrefix(language)
	return prefix + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	
	case "Spanish":
		prefix = "Hola, "
	case "French":
		prefix = "Bonjour, "
	case "German":
		prefix = "Hallo, "
	default:
		prefix = "Hello, "
	}
	return
}

func main() {
	fmt.Println(Hello("Conner", "English"))
}
