package main

import (
  "fmt"
  "log"

  "example.com/greetings"
)

func main() {
    // set props of the predefined logger,
    // including log entry prefix and a flag 
    // to disable printing the time, source title
    // and line number.
    log.SetPrefix("greetings: ")
    log.SetFlags(0)

    // A slice of names.
    names := []string{"Di", "Shu", "Na"}

    // Request greetings messages for the names.
    messages, err := greetings.Hellos(names)
    if err != nil {
        log.Fatal(err)
    }

    // If no error was returned, print the returned map
    // of messages to the console.

    fmt.Println(messages)
}

