package main

import (
	"Go/greetings"
	"fmt"
	"log"
	"time"
)

func main() {
	// set properties of the predefined Logger
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Get a greeting message and print it.
	for i := 0; i < 10; i++ {
		message, err := greetings.Hello("")
		time.Sleep(2 * time.Second)
		fmt.Println(message)

		if err != nil {
			log.Fatal(err)
			break
		}
	}
}
