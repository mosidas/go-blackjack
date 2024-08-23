package object

// Hand represents a hand of cards
type Hand struct {
	cards []Card
}

// NewHand creates a new Hand with initialized fields
func NewHand() *Hand {
	return &Hand{
		cards: []Card{},
	}
}

// Add adds a card to the hand
func (h *Hand) Add(card Card) {
	h.cards = append(h.cards, card)
}

// Hit draws one card from the deck and adds it to the hand
func (h *Hand) Hit(deck *Deck) {
	h.Add(deck.Pop())
}

// ToList returns the list of cards in the hand
func (h *Hand) ToList() []Card {
	return h.cards
}

// IsBust checks if the hand is bust (score over 21)
func (h *Hand) IsBust() bool {
	return h.GetScore() > 21
}

// GetScore calculates and returns the score of the hand
func (h *Hand) GetScore() int {
	// Calculate the score as:
	// Sum of card values less than or equal to 10 + (number of cards greater than 10) * 10
	score := 0
	aceCount := 0

	for _, card := range h.cards {
		if card.GetNumber() == 1 {
			aceCount++
			score += 1
		} else if card.GetNumber() > 10 {
			score += 10
		} else {
			score += card.GetNumber()
		}
	}

	// If there are aces, add 10 if it doesn't cause the score to exceed 21
	for aceCount > 0 && score+10 <= 21 {
		score += 10
		aceCount--
	}

	return score
}

// IsNaturalBlackjack checks if the hand is a natural blackjack (score 21 with two cards)
func (h *Hand) IsNaturalBlackjack() bool {
	return h.GetScore() == 21 && len(h.cards) == 2 && h.cards[0].GetNumber() == 1 || h.cards[1].GetNumber() == 1
}

// IsSoft17 checks if the hand is a soft 17 (score 17 with an ace)
func (h *Hand) IsSoft17() bool {
	return h.GetScore() == 17 && h.hasAce()
}

// hasAce checks if the hand contains an ace
func (h *Hand) hasAce() bool {
	for _, card := range h.cards {
		if card.GetNumber() == 1 {
			return true
		}
	}
	return false
}
