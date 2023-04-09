package main

import "fmt"

func Hello(name string) string {
	greetingString := "Hello, "
	return greetingString + name + "!"
}

func main() {
	fmt.Println(Hello("Nipun"))
}
