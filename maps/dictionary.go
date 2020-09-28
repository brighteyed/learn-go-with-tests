package maps

// Dictionary represents a key-value storage
type Dictionary map[string]string

// DictionaryErr represents an error string while operating with dictionary
type DictionaryErr string

const (
	// ErrNotFound indicates that a dictionary doesn't contain a word
	ErrNotFound = DictionaryErr("could not find a word you were looking for")

	// ErrWordExists indicates that a dictionary already has a word
	ErrWordExists = DictionaryErr("word already exists")

	// ErrWordDoesNotExist indicates that it is not possible to update a definition of a word that doesn't exist
	ErrWordDoesNotExist = DictionaryErr("cannot update a word because it doesn't exist")
)

func (e DictionaryErr) Error() string {
	return string(e)
}

// Search searches for an item in the dictionary
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

// Add adds a definition of a word to dictionary
func (d Dictionary) Add(word string, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

// Update updates a definition of the existing word
func (d Dictionary) Update(word string, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = definition
	default:
		return err
	}

	return nil
}

// Delete deletes a word from dictionary
func (d Dictionary) Delete(word string) {
	delete(d, word)
}
