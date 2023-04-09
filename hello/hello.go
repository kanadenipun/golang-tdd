package main

import "fmt"

const (
	french  = "French"
	spanish = "Spanish"

	greetingStringEnglish = "Hello, "
	greetingStringSpanish = "Hola, "
	greetingStringFrench  = "Bonjour, "
)

func Hello(name string, language string) string {

	if name == "" {
		name = "World"
	}
	return greetingPrefix(language) + name + "!"
}

func greetingPrefix(language string) (greetingText string) {
	switch language {
	case spanish:
		greetingText = greetingStringSpanish
	case french:
		greetingText = greetingStringFrench
	default:
		greetingText = greetingStringEnglish
	}
	return
}

func main() {
	fmt.Println(Hello("Nipun", ""))
}
