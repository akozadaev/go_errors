package main

import (
	"errors"
	"fmt"
)

type NetworkError struct {
	Op  string
	Err error
}

func (e *NetworkError) Error() string {
	return fmt.Sprintf("%s failed: %v", e.Op, e.Err)
}

func (e *NetworkError) Timeout() bool {
	return true // или логика определения таймаута
}

func connectToServer() error {
	return &NetworkError{
		Op:  "connect",
		Err: errors.New("connection refused"),
	}
}

func main() {
	err := connectToServer()
	if err != nil {
		var netErr *NetworkError
		if errors.As(err, &netErr) && netErr.Timeout() {
			fmt.Println("Это временная ошибка сети, повторим...")
		} else {
			fmt.Printf("Критическая ошибка: %v\n", err)
		}
	}
}
