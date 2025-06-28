package greetings

import (
	"errors"
	"fmt"
)

// Hello returns a greetings for the named person
func Hello(name string) (string, error) {
	// If no name was given, return an error with a message
	if name == "" {
		return "", errors.New("empty name")
	}
	// Return a greeting that embeds the name in a message
	message := fmt.Sprintf("Hello %v, welcome!", name)
	return message, nil
}
