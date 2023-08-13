package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "Juan"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Saludo("Juan")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Saludo("juan) = %q, %v, quiere coincidencia para %#q, nil`, msg, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	msg, err := Saludo("")
	if msg != "" || err == nil {
		t.Fatalf(`Saludo("") = %q, %v, quiere "", error`, msg, err)
	}
}
