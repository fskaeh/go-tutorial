package main

import (
	"fmt"

	greetings "example.com/greetings2"
)

func main() {
	// Get a greeting message and print it.
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}
