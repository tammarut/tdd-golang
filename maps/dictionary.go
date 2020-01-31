package maps

type Dictionary map[string]string

const (
	ErrNotFound   = DictionaryErr("couldn't find the word you were looking for")
	ErrWordExists = DictionaryErr("cannot add word because it already exists")
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

func (dic Dictionary) Update(word, definition string) {
	dic[word] = definition
}
