package main

import (
	"os"
	"reflect"
	"testing"
)

func Test_deck_toString(t *testing.T) {
	tests := []struct {
		name string
		d    deck
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.toString(); got != tt.want {
				t.Errorf("deck.toString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_deck_shuffle(t *testing.T) {
	tests := []struct {
		name string
		d    deck
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.d.shuffle()
		})
	}
}

func Test_newDeck(t *testing.T) {
	tests := []struct {
		name string
		want deck
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newDeck(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newDeck() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected %v but got %v", 52, len(d))
	}

	if d[0] != "Two of Spades" {
		t.Errorf("Expected %v but got Two Spades", d[0])
	}

}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decttesting")

	deck := newDeck()

	deck.writeToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 52 {
		t.Errorf("Expected 52 cards, got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
