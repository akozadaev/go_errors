package math

import "testing"

func TestAdd(t *testing.T) {
	result := Add(1, 4)
	if result != 5 {
		t.Errorf("Должен быть равен 5, а сейчас %f", result)
	}
	//return a + b
}

func TestDivide(t *testing.T) {
	result, err := Divide(1, 4)
	if err != nil {
		t.Fatal("Ошибка %w", err)
	}
	if result != .25 {
		t.Errorf("Должен быть равен .25, а сейчас %f", result)
	}
	//return a + b
}
