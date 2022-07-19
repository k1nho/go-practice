package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Defining slice of names to be used
	names := []string{"Kin", "John Doe", "Alice Bobber"}

	messages, err := greetings.HelloToAll(names)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messages)
}
