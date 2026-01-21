package greetings

import (
	"errors"
	"fmt"
)

func Hello(name string) (string, error) {

	// If no name was given, return an error message.

	if name == "" {
		return "", errors.New("empty name")
	}

	// Return a greeting message
	var x string
	x = "Bitch"
	message := fmt.Sprintf("Hi, %v. Welcome %v !", name, x)

	return message, nil
}
