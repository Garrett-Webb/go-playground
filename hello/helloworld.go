package main

import (
	"fmt"

	"rsc.io/quote"

	"github.com/Garrett-Webb/go-playground/greetings"
)

func main() {
	fmt.Println(quote.Go())
	// Get a greeting message and print it.
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}
