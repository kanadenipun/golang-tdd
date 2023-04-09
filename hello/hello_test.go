package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Nipun")
	want := "Hello, Nipun!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
