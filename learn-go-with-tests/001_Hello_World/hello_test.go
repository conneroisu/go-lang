package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("world", "English")
	want := "Hello, world"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "English")
		want := "Hello, Chris"
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello, world"
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("saying hello to toby", func(t *testing.T) {
		got := Hello("Toby", "English")
		want := "Hello, Toby"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello to todd", func(t *testing.T) {
		got := Hello("Todd", "English")
		want := "Hello, Todd"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hola to conner", func(t *testing.T) {
		got := Hello("Conner", "Spanish")
		want := "Hola, Conner"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hallo to conner", func(t *testing.T) {
		got := Hello("Conner", "German")
		want := "Hallo, Conner"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying bonjour to conner", func(t *testing.T) {
		got := Hello("Conner", "French")
		want := "Bonjour, Conner"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
