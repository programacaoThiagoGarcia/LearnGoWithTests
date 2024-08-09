package main

import "testing"

func TestHello(t *testing.T) {

	t.Run("Say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello to people default language", func(t *testing.T) {
		got := Hello("Thiago", "")
		want := "Hello, Thiago!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Thiago", "Spanish")
		want := "Hola, Thiago!"
		assertCorrectMessage(t, got, want)

	})

	t.Run("in Spanish empty string", func(t *testing.T) {
		got := Hello("", "Spanish")
		want := "Hola, Mundo!"
		assertCorrectMessage(t, got, want)

	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Thiago", "French")
		want := "Bonjour, Thiago!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French empty string", func(t *testing.T) {
		got := Hello("", "French")
		want := "Bonjour, Monde!"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got -> %q want -> %q", got, want)
	}
}
