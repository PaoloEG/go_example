package greetings

import (
    "fmt"
    "errors"
    "math/rand"
    "time"
)

// Hello is a function that returns a greeting for the named person
func Hello(name string) (string, error) { 
	//functions that can be called by a function not in the same package have capital letter -> exported name

    if name == "" {
        return "", errors.New("empty name")
    }
    message := fmt.Sprintf(randomFormat(), name)
    return message, nil
}
//Hellos function is the same of hello but accept an array of names as input
func Hellos(names []string) (map[string]string, error) {
    messages := make(map[string]string)
    for _, name := range names {
        message, err := Hello(name)
        if err != nil {
            return nil, err
        }
        messages[name] = message
    }
    return messages, nil
}


//init functions are executed automatically at program startup, after initializing global variables
func init() {
    rand.Seed(time.Now().UnixNano())
}

// returns one of a set of greeting messages selected at random. PRIVATE FUNCTION since it's first letter is lowercase
func randomFormat() string {
    formats := []string{
        "Hi, %v. Welcome!",
        "Great to see you, %v!",
        "Mr %v, welcome back! We missed you...",
        "Aye Aye, %v!",
    }

    // Return a randomly selected message format by specifying
    // a random index for the slice of formats.
    return formats[rand.Intn(len(formats))]
}