package object

import (
	"math/rand"
	"time"
)

// Deck represents a deck of playing cards
type Deck struct {
	deck       []Card
	deckNumber int
	rng        *rand.Rand // ローカルなランダム生成器を追加
}

// NewDeck creates a new Deck
func NewDeck() *Deck {
	return &Deck{
		deck:       []Card{},
		deckNumber: 1,
		rng:        rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

// Create creates the deck of cards and shuffles them
func (d *Deck) Create() {
	d.deck = []Card{}

	marks := []Suit{Heart, Diamond, Spade, Crab}
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}

	for i := 0; i < d.deckNumber; i++ {
		for _, mark := range marks {
			for _, number := range numbers {
				card := NewCard(mark, number)
				d.deck = append(d.deck, card)
			}
		}
	}

	// Shuffle the deck
	d.rng.Shuffle(len(d.deck), func(i, j int) {
		d.deck[i], d.deck[j] = d.deck[j], d.deck[i]
	})
}

// SetDeckNumber sets the number of decks to be used
func (d *Deck) SetDeckNumber(n int) {
	d.deckNumber = n
}

// Pop draws one card from the deck
func (d *Deck) Pop() Card {
	card := d.deck[0]
	d.deck = d.deck[1:] // Remove the top card
	return card
}

// Size returns the remaining number of cards in the deck
func (d *Deck) Size() int {
	return len(d.deck)
}
