package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
)

func openFile() error {
	_, err := os.Open("config.json")
	if err != nil {
		return fmt.Errorf("failed to open config: %w", err)
	}
	return nil
}

func loadConfig() error {
	err := openFile()
	if err != nil {

		return fmt.Errorf("loadConfig: %w", err) // ← вот здесь
	}
	return nil
}

type CustomError struct {
	Code    int
	Message string
}

func (c CustomError) Error() string  {
	return fmt.Sprintf("%d: %s", c.Code, c.Message)
}

func main() {
	//err := openFile()
	err := loadConfig()
	if err != nil {
		fmt.Println(err)
	}

	if errors.Is(err, fs.ErrNotExist) {
		fmt.Println("Config missing — using defaults")
	}
}
