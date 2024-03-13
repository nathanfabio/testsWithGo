package maps

import "errors"

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	definition, exist := d[word]

	if !exist {
		return "", errors.New("could not find the word")
	}

	return definition, nil
}