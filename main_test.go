package main

import "testing"

func TestHelloWord(t *testing.T) {
	result := "Hello World"
	expected := "Helo, World!"

	if result != expected {
		t.Errorf("got %q, want %q", result, expected)
	}
}
