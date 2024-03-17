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
	t.Run("new word", func(t *testing.T) {
		dictionary:= Dictionary{}
		word:= "test"
		definition:= "this is just a test"

		err:= dictionary.Add(word, definition)

		compareError(t, err, nil)
		compareDefinition(t, dictionary, word, definition)
	})

	t.Run("existing word", func(t *testing.T) {
		word:= "test"
		definition:= "this is just a test"
		dictionary:= Dictionary{word: definition}
		err:= dictionary.Add(word, "new test")

		compareError(t, err, ErrExistingWord)
		compareDefinition(t, dictionary, word, definition)
	})

}

func compareDefinition(t *testing.T, dictionary Dictionary, word, definition string) {
	t.Helper()

	result, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should have found added word:", err)
	}

	if definition != result {
		t.Errorf("result '%s', expected '%s'", result, definition)
	}
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word:= "teste"
		definition:= "this is just a test"
		dictionary:= Dictionary{word: definition}
		newDefinition:= "new definition"
		err:= dictionary.Update(word, newDefinition)

		compareError(t, err, nil)
		compareDefinition(t, dictionary, word, newDefinition)
	})
	
	t.Run("new word", func(t *testing.T) {
		word:= "test"
		definition:= "this is just a test"
		dictionary:= Dictionary{}

		err:= dictionary.Update(word, definition)

		compareError(t, err, ErrNonExistentWord)
	})
}

func TestDelete(t *testing.T) {
	word:= "test"
	dictionary:= Dictionary{word: "test of definition"}

	dictionary.Delete(word)

	_, err := dictionary.Search(word)
	if err != ErrNotFound {
		t.Errorf("'%s' is expected to be deleted", word)
	}
}