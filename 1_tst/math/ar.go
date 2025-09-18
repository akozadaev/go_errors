package math

import "errors"

func Add(a, b float64) float64 {
	return a + b
}

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("Ошибка деления на ноль")
	}
	return a / b, nil
}
