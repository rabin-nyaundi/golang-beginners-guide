package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// New type of deck which s a lice of strings
type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Clubs", "Hearts", "Diamonds"}
	cardValues := []string{"Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King", "Ace"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func deal(d deck, handSize int) (deck, deck) {

	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	deckString := strings.Join([]string(d), ",")
	return deckString
}

func (d deck) writeToFile(filename string) error {
	err := ioutil.WriteFile(filename, []byte(d.toString()), 0666)

	if err != nil {
		return err
	}

	fmt.Println("Done!!!!!!!!!!!")

	return err
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffle() {

	source := rand.NewSource(time.Now().Unix())
	r := rand.New(source)

	for index := range d {
		newPosition := r.Intn(len(d))
		d[index], d[newPosition] = d[newPosition], d[index]
	}
}
