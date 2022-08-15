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

    // request a greeting message
    message, err := greetings.Hello("Gladys")
    // if an error was returned, print it to the console
    // and exit program
    if err != nil {
      log.Fatal(err)
    }

    fmt.Println(message)
}

