package main

import (
	"fmt"
	"log"
	_ "log"
)

func main() {

	defer func() {
		if r := recover(); r != nil {
			log.Printf("Recovered from panic: %v", r)
		}
	}()

	//panic("not implemented")

	fmt.Println("hello")

	_, err := main2()

	fmt.Println(err)
}

func main2() (int, error) {
	var err error
	return 1, err
}
