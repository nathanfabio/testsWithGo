package main

import "testing"

func TestHello(t *testing.T) {
	checkCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper() // marks this function as a test helper
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Nathan" , "")
		want := "Hello, Nathan"
	
		checkCorrectMessage(t, got, want)
	
	})

	t.Run("'World' as a standart to an empty string", func(t *testing.T) {
		got := Hello("" , "")
		want := "Hello, World"
		
		checkCorrectMessage(t, got, want)
	})

	t.Run("in Portuguese", func(t *testing.T) {
		got := Hello("Nathan", "Portuguese")
		want := "Olá, Nathan"
	
		checkCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Nathan", "French")
		want := "Bonjour, Nathan"
	
		checkCorrectMessage(t, got, want)
	})
	
}