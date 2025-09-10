package greet

import (
	"errors"
	"fmt"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "Error", errors.New("empty name")
	}
	// return fmt.Sprintf("Hello, %s!", name)
	message := fmt.Sprintf("Hello, %s", name)
	return message, nil
}

func Bye(name string) (string, error) {
	if name == "" {
		return "Error", errors.New("empty name")
	}
	message := fmt.Sprintf("Goodbye, %s!", name)
	return message, nil
}

func Welcome(name string) (string, error) {
	if name == "" {
		return "Error", errors.New("empty name")
	}
	message := fmt.Sprintf("Welcome, %s!", name)
	return message, nil
}
