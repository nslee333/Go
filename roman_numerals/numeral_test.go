package main

import "testing"

func TestRomanNumerals(t *testing.T) {
	t.Run("returns I", func(t *testing.T) {
		got := convertToRoman(1)
		want := "I"

		assertString(t, got, want)
	})

	t.Run("", func(t *testing.T) {
		got := convertToRoman(2)
		want := "II"

		assertString(t, got, want)
	})
}

func assertString(t *testing.T, got string, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
