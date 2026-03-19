package helloworld

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got, want := Hello("Chris", ""), "Hello Chris"
		assertionCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got, want := Hello("", ""), "Hello World"
		assertionCorrectMessage(t, got, want)
	})

	t.Run("In Spanish", func(t *testing.T) {
		got, want := Hello("Elodie", "Spanish"), "Hola Elodie"
		assertionCorrectMessage(t, got, want)
	})

	t.Run("In Franch", func(t *testing.T) {
		got, want := Hello("Mike", "Franch"), "Bonjour Mike"
		assertionCorrectMessage(t, got, want)
	})
}

func assertionCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
