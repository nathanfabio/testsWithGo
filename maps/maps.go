package maps


type Dictionary map[string]string

type ErrDictionary string

const (
	ErrNotFound = ErrDictionary("couldn't find the word you're looking for")
	ErrExistingWord = ErrDictionary("you can't add the word because it already exists")
	ErrNonExistentWord = ErrDictionary("it was not possible to update the word as it does not exist")
)

func (d Dictionary) Search(word string) (string, error) {
	definition, exist := d[word]
	if !exist {
		return "", ErrNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word) 
	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrExistingWord
	default:
		return err
	}

	return nil
}

func (e ErrDictionary) Error() string {
	return string(e)
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case ErrNotFound:
		return ErrNonExistentWord
	case nil:
		d[word] = definition
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}