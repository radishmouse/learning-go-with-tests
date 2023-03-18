package iteration

import "testing"

func TestIteration(t *testing.T) {
	t.Run("tests basic Repeat", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"

		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})
	t.Run("tests basic Repeat with a different letter", func(t *testing.T) {
		repeated := Repeat("b", 5)
		expected := "bbbbb"

		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})
	t.Run("tests a different number of repetitions", func(t *testing.T) {
		repeated := Repeat("c", 10)
		expected := "cccccccccc"
		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})
	t.Run("tests with 0 passed in for number of times to repeat", func(t *testing.T) {
		repeated := Repeat("d", 0)
		expected := "ddddd"
		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
