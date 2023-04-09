package iteration

import "testing"

func TestRepeat(t *testing.T) {
	t.Run("should repeat a five times", func(t *testing.T) {
		repeated := Repeat("a")
		expected := "aaaaa"
		assertCorrectMessage(t, expected, repeated)
	})

	t.Run("should repeat b five times", func(t *testing.T) {
		repeated := Repeat("b")
		expected := "bbbbb"
		assertCorrectMessage(t, expected, repeated)
	})
}

func assertCorrectMessage(t testing.TB, expected, got string) {
	t.Helper()
	if got != expected {
		t.Errorf("expected '%v', got '%v'", expected, got)
	}
}
