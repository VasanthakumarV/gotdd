package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	assertEqualSlices := func(t *testing.T, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d for %v", got, want, numbers)
		}
	})

	t.Run("returns the sum of the given list of ints", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0, 9})
		want := []int{3, 9}

		assertEqualSlices(t, got, want)
	})

	t.Run("returns the sum of tails of the given set of slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{0, 9})
		want := []int{0, 9}

		assertEqualSlices(t, got, want)
	})
}
