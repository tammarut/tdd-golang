package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{
		"test": "this is just a test",
	}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("xxx")
		want := "couldn't find the word you were looking for"

		if err == nil {
			t.Fatal("expected to get an error")
		}

		assertStrings(t, err.Error(), want)
		assertError(t, err, ErrNotFound)
	})
}
func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		definition := "this is just a test"
		err := dictionary.Add(word, definition)

		assertError(t, err, nil)
		asssertDefinition(t, dictionary, word, definition)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictaionary := Dictionary{word: definition}
		err := dictaionary.Add(word, "new test")

		assertError(t, err, ErrWordExists)
		asssertDefinition(t, dictaionary, word, definition)
	})
}
func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		newDefinition := "new definition"

		err := dictionary.Update(word, newDefinition)

		assertError(t, err, nil)
		asssertDefinition(t, dictionary, word, newDefinition)
	})
	t.Run("new word", func(t *testing.T) {
		word := "test"
		definition := "ant"
		dictionary := Dictionary{}

		err := dictionary.Update(word, definition)

		assertError(t, err, ErrWordDoesNotExist)
	})
}
func TestDelete(t *testing.T) {
	word := "test"
	dictionary := Dictionary{word: "xxxxx"}

	dictionary.Delete(word)

	_, err := dictionary.Search(word)
	if err != ErrNotFound {
		t.Errorf("Expected %q to be deleted", word)
	}
}

func assertStrings(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
	if got == nil {
		if want == nil {
			return
		}
		t.Fatal("expected to get an error.")
	}
}
func asssertDefinition(t *testing.T, dictionary Dictionary, word, definition string) {
	t.Helper()
	got, err := dictionary.Search("test")

	if err != nil {
		t.Fatal("should find added word:", err)
	}

	if got != definition {
		t.Errorf("got %q but want %q", got, definition)
	}
}
