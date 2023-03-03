package main

import "fmt"

type EnglishBot struct{}
type SpanishBot struct{}

type Bot interface {
	getGreeting() string
}

func main() {
	eb := EnglishBot{}
	sb := SpanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b Bot) {
	fmt.Println(b.getGreeting())
}

func (eb EnglishBot) getGreeting() string {
	return "Hi There!"
}

func (sb SpanishBot) getGreeting() string {
	return "Hola!"
}
