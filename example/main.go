package main

import (
	"errors"
	"fmt"
)

// Определим эталонные (типовые) ошибки — как константы
var (
	ErrNotFound     = errors.New("not found")
	ErrInvalidInput = errors.New("invalid input")
	ErrPermission   = errors.New("permission denied")
)

// CustomError — пользовательская ошибка с контекстом
type CustomError struct {
	Code    int
	Message string
}

func (e CustomError) Error() string {
	return fmt.Sprintf("ERR_%d: %s", e.Code, e.Message)
}

// Реализуем метод Is() — чтобы errors.Is() мог работать с нашим типом
func (e CustomError) Is(target error) bool {
	switch target {
	case ErrNotFound:
		return e.Code == 404
	case ErrInvalidInput:
		return e.Code == 400
	case ErrPermission:
		return e.Code == 403
	}
	return false
}

// ValidationError — ошибка валидации
type ValidationError struct {
	Field string
	Code  string
	Value any
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation failed on field '%s': %s (value: %v)", e.Field, e.Code, e.Value)
}

// GetField — API-метод
func (e *ValidationError) GetField() string { return e.Field }

// Пример: обернём ValidationError в другую ошибку (например, через fmt.Errorf)
func validateEmail(email string) error {
	if email == "" {
		return fmt.Errorf("failed to validate email: %w", &ValidationError{Field: "email", Code: "required", Value: email})
	}
	if email != "valid@example.com" {
		return fmt.Errorf("invalid email format: %w", &ValidationError{Field: "email", Code: "invalid_format", Value: email})
	}
	return nil
}

// Пример: функция, которая может возвращать разные ошибки
func processUserInput(email string) error {
	err := validateEmail(email)
	if err != nil {
		// Обернём в системную ошибку
		return fmt.Errorf("user processing failed: %w", err)
	}
	return nil
}

func main() {
	// --- 1. ПРОВЕРКА С ПОМОЩЬЮ errors.Is() ---

	fmt.Println("=== errors.Is() ===")

	// Симулируем ошибку
	userErr := processUserInput("bad-email")
	fmt.Printf("Общая ошибка: %v\n", userErr)

	// Проверяем, является ли она ошибкой "not found" — нет
	if errors.Is(userErr, ErrNotFound) {
		fmt.Println("Это ошибка 'not found'")
	} else {
		fmt.Println("Это НЕ ошибка 'not found'")
	}

	// Проверяем, является ли она ошибкой "invalid input" — тоже нет
	if errors.Is(userErr, ErrInvalidInput) {
		fmt.Println("Это ошибка 'invalid input'")
	} else {
		fmt.Println("Это НЕ ошибка 'invalid input'")
	}

	// Проверяем, содержит ли цепочка ValidationError — ДА!
	// Но для этого нужно использовать errors.As()

	// --- 2. ПРОВЕРКА С ПОМОЩЬЮ errors.As() ---
	fmt.Println("\n=== errors.As() ===")

	var ve *ValidationError
	if errors.As(userErr, &ve) {
		fmt.Printf("Найдена ошибка валидации!\n")
		fmt.Printf("   → Поле: %s\n", ve.GetField())
		fmt.Printf("   → Код: %s\n", ve.Code)
		fmt.Printf("   → Значение: %v\n", ve.Value)
	} else {
		fmt.Println("Не найдено ValidationError в цепочке")
	}

	// --- 3. ПРОВЕРКА С ПОМОЩЬЮ errors.Is() + собственная реализация Is() ---
	fmt.Println("\n=== errors.Is() с кастомным типом ===")

	customErr := CustomError{Code: 404, Message: "Resource not found"}
	fmt.Printf("CustomError: %v\n", customErr)

	// Теперь errors.Is будет использовать наш метод Is()
	if errors.Is(customErr, ErrNotFound) {
		fmt.Println("CustomError считается ошибкой 'not found' благодаря методу Is()")
	} else {
		fmt.Println("CustomError не распознан как ErrNotFound")
	}

	if errors.Is(customErr, ErrPermission) {
		fmt.Println("Это ошибка доступа")
	} else {
		fmt.Println("Это НЕ ошибка доступа")
	}

	// --- 4. ПРИМЕР С ОШИБКОЙ, КОТОРАЯ НЕ УПАКОВАНА (прямое сравнение) ---
	fmt.Println("\n=== Прямое сравнение (без упаковки) ===")

	rawErr := &ValidationError{Field: "age", Code: "too_low", Value: 15}
	if errors.As(rawErr, &ve) {
		fmt.Printf("Прямая ошибка: поле '%s'\n", ve.GetField())
	}

	// Проверим, что errors.Is работает только с типами, у которых реализован Is()
	// В этом случае — нет, потому что rawErr не реализует Is() для ErrInvalidInput
	if errors.Is(rawErr, ErrInvalidInput) {
		fmt.Println("Ожидаемо: rawErr не имеет Is(), поэтому не сравнивается")
	} else {
		fmt.Println("Как и ожидалось: errors.Is не сработал без реализации Is()")
	}

	// --- 5. ПРИМЕР С errors.Join() (объединение нескольких ошибок) ---
	fmt.Println("\n=== errors.Join() ===")

	err1 := fmt.Errorf("first error: %w", ErrNotFound)
	err2 := fmt.Errorf("second error: %w", &ValidationError{Field: "password", Code: "too_short", Value: "123"})
	joinedErr := errors.Join(err1, err2)

	fmt.Printf("Объединённая ошибка: %v\n", joinedErr)

	// Проверим, содержит ли joinedErr ErrNotFound
	if errors.Is(joinedErr, ErrNotFound) {
		fmt.Println("errors.Is() нашёл ErrNotFound внутри объединённой ошибки!")
	}

	// Проверим, содержит ли joinedErr ValidationError
	var ve2 *ValidationError
	if errors.As(joinedErr, &ve2) {
		fmt.Printf("errors.As() нашёл ValidationError: поле '%s'\n", ve2.Field)
	} else {
		fmt.Println("Не нашёл ValidationError (может быть несколько — as берёт первую)")
	}
}
