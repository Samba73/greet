package greet

import (
	"regexp"
	"testing"
)

func TestMessageName(t *testing.T) {
	name := "Raghu"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Message("Samba")

	if !want.MatchString(msg) || err != nil {
		t.Errorf(`Message("Samba")=%q, %v, want match for %#q,nil`, msg, err, want)
	}
}

func TestMessageEmpty(t *testing.T) {
	msg, err := Message("")

	if msg != "error" || err == nil {
		t.Errorf(`Message("") = %q, %v, want "", error`, msg, err)
	}
}
