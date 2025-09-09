package greet

import "fmt"

func Hello(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}

func Bye(name string) string {
	return fmt.Sprintf("Goodbye, %s!", name)
}

func Welcome(name string) string {
	return fmt.Sprintf("Welcome, %s!", name)
}
