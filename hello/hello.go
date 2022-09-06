package main

import (
    "fmt"
    "paoloeg/greetings"
    "log"
)

func main() {
    log.SetFlags(0)
    log.SetPrefix("greetings: ")
    
    names := []string{"Mario", "Bob", "Giacomino"}

    messages, err := greetings.Hellos(names)
    if err != nil {
        log.Fatal(err)
    }
    for name := range names {
        fmt.Println((messages[names[name]]))
    }
    // fmt.Println(messages)
}