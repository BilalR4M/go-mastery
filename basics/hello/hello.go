package main

import (
	"fmt"
	"log"

	"bilal.com/greetings"
)

func main() {
	// Set properties of the predifined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Get a greeting message and print it.
	message, err := greetings.Hello("Bilze")

	messages, err := greetings.Hellos([]string{"Bilal", "Ahmed", "Saeed"})

	// If an error was returned, print it to the console
	// exit the program
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)

	fmt.Println(messages)
}
