package main

import "testing"

func TestParameter(t *testing.T) {
	got := Parameter(10.0, 10.0)
	want := 40.0

	if got != want {
		t.Errorf("Got %.2f want %.2f", got, want)
	}
}
