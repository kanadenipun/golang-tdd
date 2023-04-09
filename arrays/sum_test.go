package arrays

import "testing"

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
		var gotArray [2]int
		copy(gotArray[:], got)
		wantArray := [2]int{3, 9}
		if wantArray != gotArray {
			t.Errorf("expected '%v', got '%v'", wantArray, gotArray)
		}

	})

}
