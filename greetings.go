package golang_greetings

import (
	"errors"
	"fmt"
	"time"
	"math/rand"
	"strings"
)

func GreetPatterns(liked bool) []string {
	if (liked) {
		return []string{
			"Hello %v, nice to meet you!",
			"Welcome %v",
			"%v! Great to see ya",
		}
	}

	return []string{
		"Oh great, it's %v... NOT",
		"%v is here, that means the fun just left",
		"Boo, you stink%.0v",
	}
}

// This sets the seed to be precise to the program's start time
func init() {
	rand.Seed(time.Now().UnixNano())
}

// Return a greeting for a name based on if the entity is liked
func Greet(name string, liked bool) string {
	patterns := GreetPatterns(liked)
	return fmt.Sprintf(patterns[rand.Intn(len(patterns))], name)
}

// returns a greeting for the named person
func Hello(name string) (string, error) {
	// if the name is invalid (empty) give an error
	if (name == "") {
		return "", errors.New("empty name")
	}

	if (strings.ToLower(name) == "evan") {
		return Greet(name, false), nil
	}
	return Greet(name, true), nil
}