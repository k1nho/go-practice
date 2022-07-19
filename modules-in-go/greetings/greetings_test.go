package greetings

import (
	"regexp"
	"testing"
)

func TestHello(t *testing.T) {
	name := "Kin"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("Kin")

	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Kin") = %q %v want match for %#q, nil`, msg, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q %v, want "", error`, msg, err)
	}
}
