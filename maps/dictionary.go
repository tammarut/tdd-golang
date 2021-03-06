package maps

type Dictionary map[string]string

const (
	ErrNotFound         = DictionaryErr("couldn't find the word you were looking for")
	ErrWordExists       = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("cannot update word because it does not exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

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

func (dic Dictionary) Update(word, definition string) error {
	_, err := dic.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		dic[word] = definition
	default:
		return err
	}

	return nil
}

func (dic Dictionary) Delete(word string) {
	delete(dic, word)
}
