package object

import "strconv"

type Card struct {
	Suit string
	Rank int
}

func NewCard(suit string, rank int) Card {
	return Card{Suit: suit, Rank: rank}
}

func (c *Card) Score() int {
	if c.Rank >= 10 {
		return 10
	} else {
		return c.Rank
	}
}

func (c *Card) String() string {
	return c.Suit + ":" + strconv.Itoa(c.Rank)
}
