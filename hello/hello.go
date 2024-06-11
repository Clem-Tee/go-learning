package main

import (
	"fmt"

	"rsc.io/quote"

	"example.com/greetings"
)

func main() {

	// Get a greeting message and print it.
	message := greetings.Hello("Gladys")
	fmt.Println(message)

	fmt.Println(quote.Go())
}
