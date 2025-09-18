package main

import (
	"errors"
	"fmt"
)

var ErrPermissionDenied = errors.New("permission denied")

func openFile(filename string) error {
	if filename == "" {
		return fmt.Errorf("filename is empty")
	}
	if filename == "secret.txt" {
		return fmt.Errorf("access to %s: %w", filename, ErrPermissionDenied)
	}
	return nil
}

func main() {
	err := openFile("secret.txt")
	if err != nil {
		if errors.Is(err, ErrPermissionDenied) {
			fmt.Println("Доступ запрещён!")
		} else {
			fmt.Printf("Ошибка: %v\n", err)
		}
	}
}
