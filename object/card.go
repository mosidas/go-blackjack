package object

import "fmt"

// Suit represents the suit of the card
type Suit int

const (
	Heart Suit = iota
	Crab
	Spade
	Diamond
)

// Stringerを使ってSuit型の文字列表現を生成
// このコメントはstringerツールが使われることを示すために残します。
// go:generate stringer -type=Suit

// Card represents a playing card
type Card struct {
	suit   Suit
	number int
}

// NewCard creates a new Card with the given suit and number
func NewCard(suit Suit, number int) Card {
	return Card{suit: suit, number: number}
}

// GetText returns the textual representation of the card (e.g., "♡K", "♣4")
func (c Card) GetText() string {
	var mark string
	switch c.suit {
	case Heart:
		mark = "♡"
	case Crab:
		mark = "♣"
	case Spade:
		mark = "♠"
	case Diamond:
		mark = "♦"
	default:
		mark = ""
	}

	var noStr string
	switch c.number {
	case 1:
		noStr = "A"
	case 11:
		noStr = "J"
	case 12:
		noStr = "Q"
	case 13:
		noStr = "K"
	default:
		noStr = fmt.Sprintf("%d", c.number)
	}

	return mark + noStr
}

// GetNumber returns the number of the card
func (c Card) GetNumber() int {
	return c.number
}
