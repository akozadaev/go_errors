package math

import (
	"testing"
)

func TestDivide2(t *testing.T) {
	tests := []struct {
		name      string
		a, b      float64
		expected  float64
		expectErr bool
	}{
		{"normal division", 10, 2, 5, false},
		{"division by zero", -1, -0, 0, true},
		{"negative", -1, -1, 1, false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := Divide(test.a, test.b)
			if (err != nil) != test.expectErr {
				t.Fatal("Ошибка %v, %v", test.expectErr, err)
			}
			if !test.expectErr && result != test.expected {
				t.Errorf("должно быть %f, получено %f", test.expected, result)
			}
		})
	}
}
