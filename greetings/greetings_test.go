package greetings

import (
	"regexp"
	"testing"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestHelloName(t *testing.T) {
    name := "Gladys"
    want := regexp.MustCompile(`\b`+name+`\b`)
    msg, err := Hello("Gladys")
    if !want.MatchString(msg) || err != nil {
        t.Fatalf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
    }
}

// TestHelloEmpty calls greetings.Hello with an empty string,
// checking for an error.
func TestHelloEmpty(t *testing.T) {
    msg, err := Hello("")
    if msg != "" || err == nil {
        t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
    }
}

// TestHelloNames calls greetings.Hellos
func TestHelloNames(t *testing.T) {
	names := []string{"test","test2","test3"}
    msgs, err := Hellos(names)
	if err != nil {
		t.Fatalf("Hellos function returned an error")
	}
    for _, name := range names {
		msg, ok := msgs[name]
		matched, _ := regexp.MatchString(`\b`+name+`\b`, msg)
		if !ok || !matched {
			t.Fatalf(`Hellos produce an error on name %s. Received nil`,name)
		}
	}
}