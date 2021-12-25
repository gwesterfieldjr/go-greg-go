package main

import "fmt"

type bot interface {
	getGreeting() string
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

type englishBot struct{}
type spanishBot struct{}

// omit reciever value if not using it
func (englishBot) getGreeting() string {
	return "Hello"
}

// omit reciever value if not using it
func (spanishBot) getGreeting() string {
	return "Holla"
}

func main() {
	englishBot := englishBot{}
	spanishBot := spanishBot{}

	printGreeting((englishBot))
	printGreeting((spanishBot))
}
