package helloworld

import "fmt"

// Create a constant name and prefix for each language
const (
	spanish = "Spanish"
	french  = "French"
	german  = "German"
	english = "English"

	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
	germanHelloPrefix  = "Hallo, "
	englishHelloPrefix = "Hello, "
)

// Hello returns a greeting in the specified language
func Hello(name, language string) string {
	if name == "" {
		name = "world"
	}
	prefix := greetingPrefix(language)
	return prefix + name
}

// greetingPrefix returns the appropriate prefix for the specified language
func greetingPrefix(language string) (prefix string) {
	switch language {
	case "Spanish":
		prefix = spanishHelloPrefix
	case "French":
		prefix = frenchHelloPrefix
	case "German":
		prefix = germanHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

// main function to test the Hello function
func main() {
	fmt.Println(Hello("Conner", "English"))
	fmt.Println(Hello("Conner", "Spanish"))
	fmt.Println(Hello("Conner", "French"))
	fmt.Println(Hello("Conner", "German"))
}
