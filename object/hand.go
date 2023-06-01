package object

type Hand struct {
	Cards []Card
}

func NewHand() *Hand {
	return &Hand{}
}

func (h *Hand) Add(card *Card) {
	h.Cards = append(h.Cards, *card)
}

func (h *Hand) Score() int {
	score := 0
	aceExists := false
	for _, card := range h.Cards {
		score += card.Score()
		//Ace
		if card.Rank == 1 {
			aceExists = true
		}
	}

	//Ace
	if aceExists && score <= 11 {
		score += 10
	}

	return score
}
