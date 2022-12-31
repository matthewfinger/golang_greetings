package greetings

import "fmt"

// Return a greeting for a name based on if the entity is liked
func Greet(name string, liked bool) string {
	if (liked) {
		return fmt.Sprintf("Hello %v, nice to meet you!", name)
	} else {
		return fmt.Sprintf("Oh, great, it's %v", name)
	}
}

// returns a greeting for the named person
func Hello(name string) string {
	return Greet(name, true)
}