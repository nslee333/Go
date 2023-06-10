package main

import "testing"

// Helpers
func assertStrings(t testing.TB, got, want, given string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q given %q", got, want, given)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

// Tests
func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("Known word", func(t *testing.T) {

		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want, "test")
	})

	t.Run("Unknown word", func(t *testing.T) {
		_, got := dictionary.Search("unknown")

		assertError(t, got, ErrNotFound)
	})

}

func TestAdd(t *testing.T) {
	dictionary := Dictionary{}

	dictionary.Add("test", "this is just a test")
	want := "this is just a test"

	got, err := dictionary.Search("test")

	if err != nil {
		t.Fatal("Should find word:", err)
	}

	assertStrings(t, got, want, "this is not a test")
}
