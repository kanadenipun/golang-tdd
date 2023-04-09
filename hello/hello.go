package main

import "fmt"

func Hello(name string) string {
	greetingString := "Hello, "
	if name == "" {
		name = "World"
	}
	return greetingString + name + "!"
}

func main() {
	fmt.Println(Hello("Nipun"))
}
