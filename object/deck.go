package object

import "math/rand"

type Deck struct {
	Cards []Card
}

func NewDeck() Deck {
	deck := Deck{}
	for _, suit := range []string{"Spade", "Heart", "Diamond", "Club"} {
		for rank := 1; rank <= 13; rank++ {
			deck.Cards = append(deck.Cards, NewCard(suit, rank))
		}
	}

	// Shuffle the deck
	for i := 0; i < len(deck.Cards); i++ {
		j := rand.Intn(len(deck.Cards))
		deck.Cards[i], deck.Cards[j] = deck.Cards[j], deck.Cards[i]
	}

	return deck
}

func (d *Deck) Draw() Card {
	card := d.Cards[0]
	d.Cards = d.Cards[1:]
	return card
}
