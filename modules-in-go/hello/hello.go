package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	message := greetings.Hello("Kin")
	fmt.Println(message)
}
