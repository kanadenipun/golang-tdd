package main

import "fmt"

func Hello(name string, lang string) string {
	greetingStringEnglish := "Hello, "
	greetingStringSpanish := "Hola, "
	greetingStringFrench := "Bonjour, "
	var greetingText string
	if name == "" {
		name = "World"
	}

	switch lang {
	case "Spanish":
		greetingText = greetingStringSpanish + name + "!"
	case "French":
		greetingText = greetingStringFrench + name + "!"
	default:
		greetingText = greetingStringEnglish + name + "!"
	}
	return greetingText
}

func main() {
	fmt.Println(Hello("Nipun", ""))
}
