package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello returns a greetings for the named person
func Hello(name string) (string, error) {
	// If no name was given, return an error with a message
	if name == "" {
		return "", errors.New("empty name")
	}
	// Return a greeting that embeds the name in a message
	message := fmt.Sprintf(randomFormat(), name)
	// message := fmt.Sprint(randomFormat()) <- Error test
	return message, nil
}

// Hellos returns a map that associates each of the named people
// with a greeting message
func Hellos(names []string) (map[string]string, error) {
	// A map tp asspcoate names with messages
	messages := make(map[string]string)
	// Loop through the recived slice of names, calling
	// the Hello function to get a message for each name
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		// In the map, associate the retrieved message with
		// the name
		messages[name] = message
	}
	return messages, nil
}

// randomFormat returns one of a set of greeting message. The returned
// message is selected at random
func randomFormat() string {
	// A slice of message formats
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v",
		"Hail %v, well met!",
	}

	return formats[rand.Intn(len(formats))]
}
