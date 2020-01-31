package maps

import "errors"

type Dictionary map[string]string

func (dic Dictionary) Search(word string) (string, error) {
	defination, ok := dic[word]
	if !ok {
		return "", errors.New("couldn't find the word you were looking for")
	}
	return defination, nil
}
