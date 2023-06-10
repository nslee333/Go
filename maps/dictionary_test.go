package main

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("Known word", func(t *testing.T) {

		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want, "test")
	})

	t.Run("Unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		want := "could not find the word you're looking for"

		if err == nil {
			t.Fatal("expected to get an error")
		}

		assertStrings(t, err.Error(), want, "unknown")
	})

}

func assertStrings(t testing.TB, got, want, given string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q given %q", got, want, given)
	}
}
