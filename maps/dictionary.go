package maps

import "errors"

type Dictionary map[string]string

var ErrNotFound = errors.New("couldn't find the word you were looking for")

func (dic Dictionary) Search(word string) (string, error) {
	defination, ok := dic[word]
	if !ok {
		return "", ErrNotFound
	}
	return defination, nil
}
