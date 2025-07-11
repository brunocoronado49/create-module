package main

import (
	"brunocoronado49/greetings"
	"fmt"
	"log"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and flag to disable printing
	// the time, source file and line number
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// A slice of names
	names := []string{"Yor", "Melina", "Albedo"}

	// Request greeting message for the names
	messages, err := greetings.Hellos(names)

	// if an error was returned, print it to the console and
	// exit the program
	if err != nil {
		log.Fatal(err)
	}

	// if no error was returned, print the returned map of
	// messages to the console
	fmt.Println(messages)
}
