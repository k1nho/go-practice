package greetings

import (
	"errors"
	"fmt"
)

func Hello(name string) (string, error) {

	if name == "" {
		return "", errors.New("name was empty")
	}

	message := fmt.Sprintf("Hi, %v Welcome to Go!", name)
	return message, nil
}
