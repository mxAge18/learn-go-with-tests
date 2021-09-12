package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"where": "I am living in china now"}
	t.Run("Konwn words search", func(t *testing.T) {
		got, _ := dictionary.Search("where")
		want := "I am living in china now"
		assertStrings(t, got, want)
	})
	t.Run("unkonwn words search", func(t *testing.T) {

		_, got := dictionary.Search("unknown")
		
		assertError(t, got, ErrorNotFound)
	})

}

func TestAdd(t *testing.T) {

	dictionary := Dictionary{}
	t.Run("a new word", func(t *testing.T) {
		key := "test"
		val := "this is a test"
		err := dictionary.Add(key, val)
		assertError(t, err, nil)
		assertDefination(t, dictionary, key, val)
	})
	t.Run("add existing word", func(t *testing.T) {
		key := "exist"
		val := "this is a exist test"
		dictionary := Dictionary{key: val}
		err := dictionary.Add(key, val)
		assertError(t, err, ErrorWordExist)
		assertDefination(t, dictionary, key, val)
	})
}

func TestUpdate(t *testing.T) {

	t.Run("test update exist val", func(t *testing.T) {
		word :="test"
		defination := "old val"
		dic := Dictionary{word:defination}
		newDefination := "new val"
		err := dic.Update(word, newDefination)
		assertError(t, err, nil)
		assertDefination(t, dic, word, newDefination)
	})
	t.Run("update word not exist", func(t *testing.T) {
		word :="test"
		defination := "old val"
		dic := Dictionary{word:defination}
		notExistWord := "newTest"
		newDefination := "new val"
		err := dic.Update(notExistWord, newDefination)
		assertError(t, err, ErrNonExistWord)
		assertDefination(t, dic, word, defination)
	})
}

func TestDelete(t *testing.T) {
	t.Run("delete exist word", func(t *testing.T) {
		word :="test"
		defination := "old val"
		dic := Dictionary{word:defination}
		dic.Delete(word)
	})
	t.Run("delete not exist word", func(t *testing.T) {
		word :="test"
		defination := "old val"
		notExistWord := "maxu"
		dic := Dictionary{word:defination}
		dic.Delete(notExistWord)
	})
}
func assertStrings(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got '%s' want '%s' given '%s'", got, want, "where")
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()

    if got != want {
        t.Errorf("got error '%s' want '%s'", got, want)
    }
}

func assertDefination(t *testing.T, dic Dictionary, key, val string) {
	t.Helper()
	got ,err := dic.Search(key)
	if err != nil {
		t.Fatal("should find the word:", err)
	}

	if got != val {
		t.Errorf("got error '%s' want '%s'", got, val)
	}
}