package maps

import "errors"

type Dictionary map[string]string

var (
	ErrNotFound   = errors.New("couldn't find the word you were looking for")
	ErrWordExists = errors.New("cannot add word because it already exists")
)

func (dic Dictionary) Search(word string) (string, error) {
	defination, ok := dic[word]
	if !ok {
		return "", ErrNotFound
	}
	return defination, nil
}

func (dic Dictionary) Add(word, defination string) error {
	_, err := dic.Search(word)
	switch err {
	case ErrNotFound:
		dic[word] = defination
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}
