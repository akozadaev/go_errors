package main

import (
	"errors"
	"fmt"
)

type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation failed on field '%s': %s", e.Field, e.Message)
}

func validateAge(age int) error {
	if age < 0 || age > 150 {
		return &ValidationError{
			Field:   "age",
			Message: "must be between 0 and 150",
		}
	}
	return nil
}

func main() {
	err := validateAge(-5)
	if err != nil {
		var validationErr *ValidationError
		if errors.As(err, &validationErr) {
			fmt.Printf("Ошибка валидации в поле: %s, сообщение: %s\n",
				validationErr.Field, validationErr.Message)
		} else {
			fmt.Printf("Обычная ошибка: %v\n", err)
		}
	}
}
