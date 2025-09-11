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
	// message := fmt.Sprint(randomFormat()) // to test the code
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

func Messages(names []string) (map[string]string, error) {
	messages := make(map[string]string)

	for _, name := range names {
		message, err := Message(name)

		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
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
