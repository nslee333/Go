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
	// ! Black box:
	// ! Takes in: any number of int slices
	// ! Returns: 1 slice, with each element being the sum of each argument slice.
	// * ex: func([]int{0, 3} []int{0,1}) => []int{3, 1}

	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Got %v want %v", got, want)
	}
}
