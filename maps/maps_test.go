package maps

import "testing"


func TestSearch(t *testing.T) {
	dictionary:= Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		result, _ := dictionary.Search("test")
		expected:= "this is just a test"

		compareStrings(t, result, expected)		
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")

		compareError(t, err, ErrNotFound)
	})
}

func compareStrings(t *testing.T, result, expected string) {
	t.Helper()

	if result != expected {
		t.Errorf("result '%s', expected '%s', given '%s'", result, expected, "test")
	}
}

func compareError(t *testing.T, result, expected error) {
	t.Helper()

	if result != expected {
		t.Errorf("result error '%s', expected '%s'", result, expected)
	}
}

func TestAdd(t *testing.T) {
	dictionary:= Dictionary{}
	dictionary.Add("test", "this is just a test")

	expected:= "this is just a test"
	result, err := dictionary.Search("test")
	if err != nil {
		t.Fatal("couldn't find the word added")
	}


	if expected != result {
		t.Errorf("result '%s', expected '%s'", result, expected)
	}
}