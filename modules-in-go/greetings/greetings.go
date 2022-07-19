package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func Hello(name string) (string, error) {

	if name == "" {
		return "", errors.New("name was empty")
	}

	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
	// slice (similar to vector in c++)
	formats := []string{
		"Hi, %v Welcome to go",
		"Great to see you, %v",
		"YOYOYO its your one and only %v",
	}

	return formats[rand.Intn(len(formats))]
}
