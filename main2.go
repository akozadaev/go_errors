package main

import (
	"fmt"
)

type CustomError struct {
	Code    int
	Message string
}

func (c *CustomError) Error() string {
	return fmt.Sprintf("%d: %s", c.Code, c.Message)
}

func main() {
	// 1
	customError := CustomError{400, "custom error"}
	fmt.Println(customError)
	// 2
	valErr := NewValidationError("Phone", "400", "неверный номер телефона", map[string]string{"error": "not number"})
	//fmt.Println(valErr)
	fmt.Println(valErr.Error())
	fmt.Println(valErr.Code)
	fmt.Println(valErr.Field)
}

type ValidationError struct {
	Field   string
	Code    string
	Value   any
	Details map[string]string
}

func (v *ValidationError) Error() string {
	return fmt.Sprintf("%s: %s: %s", v.Field, v.Code, v.Value)
}

func (v *ValidationError) IsFieldError() bool {
	return true
}
func (v *ValidationError) GetField() string {
	return v.Field
}

func NewValidationError(field, code string, value any, details map[string]string) *ValidationError {
	return &ValidationError{field, code, value, details}
}
