package iteration

import "testing"

func TestRepeat(t *testing.T) {
	t.Run("should repeat a five times", func(t *testing.T) {
		repeated := Repeat("a")
		expected := "aaaaa"
		if repeated != expected {
			t.Errorf("expected '%v', got '%v'", expected, repeated)
		}
	})
}
