package math

import (
	"testing"
)

func TestAdd2(t *testing.T) {
	tests := []struct {
		name     string
		a, b     float64
		expected float64
	}{
		{"positive", 2, 3, 5},
		{"negative", -1, -1, -2},
		{"mixed", -1, 1, 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := Add(test.a, test.b)
			if result != test.expected {
				t.Errorf("должно быть %f, получено %f", test.expected, result)
			}
		})
	}
}
