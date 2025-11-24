package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Returns a greeting
func Hello(name string) (string, error) {
	// if no name was given return an error with a message.
	if name == "" {
		return "", errors.New("Empty name")
	}

	// if a name is given return the value
	message := fmt.Sprintf(RandomFormat(), name)
	return message, nil
}

// Random format
func RandomFormat() string {
	// A slice of messages
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you %v",
		"Hola %v, como te va",
	}

	return formats[rand.Intn(len(formats))]
}
