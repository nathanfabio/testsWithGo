package maps

import "errors"

type Dictionary map[string]string

var ErrNotFound = errors.New("couldn't find the word you're looking for")

func (d Dictionary) Search(word string) (string, error) {
	definition, exist := d[word]
	if !exist {
		return "", ErrNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(word, definition string) {
	d[word] = definition
}