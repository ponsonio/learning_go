package main

import "testing"

func TestNewDeck(t *testing.T) {

	d := newDeck()

	if len(d) != 16 {
		t.Errorf("worong number of cards, got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("1st card Should be Ace of Spades, is %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("last card Should be Four of Clubs, is %v", d[len(d)-1])
	}

}
∫