package math

import "testing"

func TestMath(t *testing.T) {
	t.Run("Add", func(t *testing.T) {
		t.Parallel()
		result := Add(1, 4)
		if result != 5 {
			t.Errorf("Должен быть равен 5, а сейчас %f", result)
		}
	})

	t.Run("Divide", func(t *testing.T) {
		t.Parallel()
		result, err := Divide(1, 4)
		if err != nil {
			t.Fatal("Ошибка %w", err)
		}
		if result != .25 {
			t.Errorf("Должен быть равен .25, а сейчас %f", result)
		}
	})
}
