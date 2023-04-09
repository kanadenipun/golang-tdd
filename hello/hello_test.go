package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Say Hello when a name is passed", func(t *testing.T) {
		got := Hello("Nipun", "")
		want := "Hello, Nipun!"

		assertCorrectMessage(t, got, want)
	})
	t.Run("Say Hello World when name is not passed", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World!"

		assertCorrectMessage(t, got, want)
	})
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Nipun", "Spanish")
		want := "Hola, Nipun!"

		assertCorrectMessage(t, got, want)
	})
	t.Run("in French", func(t *testing.T) {
		got := Hello("Nipun", "French")
		want := "Bonjour, Nipun!"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
