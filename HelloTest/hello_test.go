package main

import "testing"

func TestHello(t *testing.T) {

	t.Run("Hello, Chris", func(t *testing.T) {
		want := "Hello, Chris"
		got := Hello("Chris", "")
		assertCorrectMessage(t, want, got)
	})

	t.Run("Hello, World", func(t *testing.T) {
		want := "Hello, World"
		got := Hello("","")
		assertCorrectMessage(t, want, got)
	})

	t.Run("Greetings on Spanish", func(t *testing.T) {
		want:="Holla, Elodie"
		got:= Hello("Elodie", "Spanish")
		assertCorrectMessage(t, want, got)
	})
}

func assertCorrectMessage(t testing.TB, want, got string) {
	t.Helper()
	if want != got {
		t.Errorf("want %q get %q", want, got)
	}
}
