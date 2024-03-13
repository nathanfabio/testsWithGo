package maps


import "testing"

func TestSearch(t *testing.T) {
	dictionary:= Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
	result, _:= dictionary.Search("test")
	expected:= "this is just a test"

	compareStrings(t, result, expected)
	})


	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")

		if err == nil {
			t.Fatal("an error is expected")
		}
	})
}

func compareStrings(t *testing.T, result, expected string) {
	t.Helper()

	if result != expected {
		t.Errorf("result '%s', expected '%s', given '%s'", result, expected, "test")
	}
}