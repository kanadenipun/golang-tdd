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
	if lang == "Spanish" {
		greetingText = greetingStringSpanish + name + "!"
	} else if lang == "French" {
		greetingText = greetingStringFrench + name + "!"
	} else {
		greetingText = greetingStringEnglish + name + "!"
	}
	return greetingText
}

func main() {
	fmt.Println(Hello("Nipun", ""))
}
