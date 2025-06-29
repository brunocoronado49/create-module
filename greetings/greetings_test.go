package greetings

import (
	"regexp"
	"testing"
)

// TestHelloName calls greetings.Hello with a name, cheking
// for a valid return value
func TestHelloName(t *testing.T) {
	name := "Yor Forger"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("Yor Forger")
	if !want.MatchString(msg) || err != nil {
		t.Errorf(`Hello("Yor Forger") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

// TestHelloEmpty calls greetings.Hello with an empty string,
// cheking for an error
func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Errorf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
