package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Say Hello when a name is passed", func(t *testing.T) {
		got := Hello("Nipun")
		want := "Hello, Nipun!"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("Say Hello World when name is not passed", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World!"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
