package main

import "fmt"

func Hello(name string) string {
	greetingString := "Hello, "
	defaultGreetingString := greetingString + "World!"
	if len(name) < 1 {
		return defaultGreetingString
	}
	return greetingString + name + "!"
}

func main() {
	fmt.Println(Hello("Nipun"))
}
