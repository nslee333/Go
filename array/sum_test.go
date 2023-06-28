package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	numbers := []int{1, 2, 3, 4, 5}

	got := Sum(numbers)
	want := 15

	if got != want {
		t.Errorf("Got %v want %v gave %v", got, want, numbers)
	}
}

func TestSumAll(t *testing.T) {

	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Got %v want %v", got, want)
	}

	checkSums(t, got, want)
}

func TestSumAllTails(t *testing.T) {
	t.Run("Make sums of the slices", func(t *testing.T) {
		numbersToSum1 := []int{1, 2, 3, 4, 5}
		numbersToSum2 := []int{1, 2, 3, 4, 5}
		numbersToSum3 := []int{1, 2, 3, 4, 5}

		got := SumAllTails(numbersToSum1, numbersToSum2, numbersToSum3)
		want := []int{14, 14, 14}

		checkSums(t, got, want)
	})

	t.Run("Safely sum empty slices", func(t *testing.T) {

		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}

		checkSums(t, got, want)
	})

}

func checkSums(t *testing.T, got, want []int) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Got %v want %v", got, want)
	}
}

func TestReduce(t *testing.T) {
	t.Run("multiplication of all elements", func(t *testing.T) {
		multiply := func(x, y int) int {
			return x * y
		}

		AssertEqual(t, Reduce([]int{1, 2, 3}, multiply, 1), 6)
	})

	t.Run("concatenate strings", func(t *testing.T) {
		concatenate := func(x, y string) string {
			return x + y
		}

		AssertEqual(t, Reduce([]string{"a", "b", "c"}, concatenate, ""), "abc")
	})
}

func AssertEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func AssertNotEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got == want {
		t.Errorf("didn't want %v", got)
	}
}

func AssertTrue(t *testing.T, got bool) {
	t.Helper()
	if !got {
		t.Errorf("got %v, want true", got)
	}
}

func AssertFalse(t *testing.T, got bool) {
	t.Helper()
	if got {
		t.Errorf("got %v, want false", got)
	}
}
