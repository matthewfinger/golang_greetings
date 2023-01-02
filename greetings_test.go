package golang_greetings

import (
	"testing"
	"regexp"
	"fmt"
)

// Make sure Hello returns a greeting specific to the name (note this will fail if the test name is Evan)
func TestHelloName(t *testing.T) {
	names := []string{"Gladys", "Marty", "matt", "#$T#GRSDV",}
	perName := 10
	total := perName * len(names)
	complete := 0
	for _, name := range names {
		pattern := regexp.MustCompile(`(?:^|\b|\s)` + regexp.QuoteMeta(name) + `(?:\b|\s|$)`)

		// since it gets a random name, lets try over and over again
		for i := 0; i < perName; i++ {
			message, err := Hello(name)
			if !pattern.MatchString(message) || err != nil {
				t.Fatalf(`Hello(%#v) = %q, %v, want match for %#q, nil`, name, message, err, pattern)
				return
			}
			complete++
			fmt.Printf("\rPassed %v of %v tests of Hello\t\t", complete, total)
		}
	}
	fmt.Println("")
}

func TestEmptyName(t *testing.T) {
	message, err := Hello("")
	if (message != "" || err == nil) {
		t.Fatalf(`Hello("") resulted in %#v, %v, expected "", error`, message, err)
	}
}