package main

// type deck []string

func main() {
	cards := newDeckFromFile("cards.txt")
	cards.shuffle()
	cards.print()
}
