package integers

import "testing"

func TestAdders(t *testing.T) {
	t.Run("check 2 plus 2 should return 4", func(t *testing.T) {
		sum := Add(2, 2)
		expected := 4
		checkCorrectSum(t, sum, expected)
	})
	t.Run("check 2 plus 3 should return 5", func(t *testing.T) {
		sum := Add(2, 3)
		expected := 5
		checkCorrectSum(t, sum, expected)
	})
}

func checkCorrectSum(t testing.TB, sum int, expected int) {
	t.Helper()
	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}
