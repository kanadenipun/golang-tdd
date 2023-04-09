package arrays

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("expected '%v', got '%v'", want, got)
		}
	})
}

func TestSumAll(t *testing.T) {
	t.Run("test sum of collection of collection of numbers", func(t *testing.T) {

		var got = SumAll([]int{1, 2}, []int{0, 9})
		want := []int{3, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("expected '%v', got '%v'", want, got)
		}
	})
}

func TestSumAllTrails(t *testing.T) {

	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()

	}
	t.Run("test sum of all slice trails", func(t *testing.T) {

		var got = SumAllTrails([]int{1, 2, 3}, []int{0, 5, 5})
		want := []int{5, 10}
		checkSums(t, got, want)
	})
	t.Run("safely test sum of all slice trails", func(t *testing.T) {

		var got = SumAllTrails([]int{1, 2, 3}, []int{0, 5, 5}, []int{1})
		want := []int{5, 10, 0}
		checkSums(t, got, want)
	})
}
