package main

import (
	"errors"
	"fmt"
)

var ErrNotFound = errors.New("not found")

func findUser(id int) error {
	if id <= 0 {
		return fmt.Errorf("invalid id: %d", id)
	}
	if id == 999 {
		return fmt.Errorf("user not found: %w", ErrNotFound) // оборачиваем
	}
	return nil
}

func main() {
	err := findUser(999)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			fmt.Println("Обработали ошибку: пользователь не найден")
		} else {
			fmt.Printf("Другая ошибка: %v\n", err)
		}
	}
}
