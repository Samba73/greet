package greet

import (
	"errors"
	"fmt"
	"math/rand"
)

func Message(name string) (string, error) {
	if name == "" {
		return "error", errors.New("empty name")
	}
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func randomFormat() string {
	formats := []string{
		"Hi, %v, Welcome!",
		"Great to see you, %v",
		"Hail, %v! Well met!",
	}

	return formats[rand.Intn(len(formats))]
}

// func Hello(name string) (string, error) {
// 	if name == "" {
// 		return "Error", errors.New("empty name")
// 	}
// 	// return fmt.Sprintf("Hello, %s!", name)
// 	message := fmt.Sprintf("Hello, %s", name)
// 	return message, nil
// }

// func Bye(name string) (string, error) {
// 	if name == "" {
// 		return "Error", errors.New("empty name")
// 	}
// 	message := fmt.Sprintf("Goodbye, %s!", name)
// 	return message, nil
// }

// func Welcome(name string) (string, error) {
// 	if name == "" {
// 		return "Error", errors.New("empty name")
// 	}
// 	message := fmt.Sprintf("Welcome, %s!", name)
// 	return message, nil
// }
