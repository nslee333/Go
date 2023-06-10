package main

import "testing"

func TestParameter(t *testing.T) {
	got := Parameter(10.0, 10.0)
	want := 40.0

	assertCorrectResponse(t, got, want)
}

func TestArea(t *testing.T) {
	got := Area(10.0, 10.0)
	want := 100.0

	assertCorrectResponse(t, got, want)
}

func assertCorrectResponse(t *testing.T, got, want float64) {
	t.Helper()

	if got != want {
		t.Errorf("Got %.2f want %.2f", got, want)
	}
}
