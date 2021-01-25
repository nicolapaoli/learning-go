package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
	"example.com/greetings"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Request a greeting message.
	message, err := greetings.Hello(randomName())


	// If an error was returned, print it to the console and
	// exit the program

	if err != nil {
		log.Fatal(err)
	}
	
	// If no error was returned, print the returned message
	// to the console
	fmt.Println(message)

	greatAll()
}

func greatAll() {
	// Creating a slice of names
	names := []string{"Nic", "Milli", "Biga", "Fred"}
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messages)
}

// Initialising Random
func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomName() string {
	names := []string {
		"Nicola", "Milena", "Damiano", "",
	}
	return names[rand.Intn(len(names))]
}