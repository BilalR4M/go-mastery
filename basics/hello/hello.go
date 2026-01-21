package main

import (
	"fmt"

	"bilal.com/greetings"
)

func main() {
	// Get a greeting message and print it.
	message, err := greetings.Hello("")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(message)
}
