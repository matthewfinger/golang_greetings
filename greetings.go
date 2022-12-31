package golang_greetings

import (
	"errors"
	"fmt"
)

// Return a greeting for a name based on if the entity is liked
func Greet(name string, liked bool) string {
	if (liked) {
		return fmt.Sprintf("Hello %v, nice to meet you!", name)
	} else {
		return fmt.Sprintf("Oh, great, it's %v", name)
	}
}

// returns a greeting for the named person
func Hello(name string) (string, error) {
	// if the name is invalid (empty) give an error
	if (name == "") {
		return "", errors.New("empty name")
	}
	return Greet(name, true), nil
}