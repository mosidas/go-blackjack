package object

// Player represents a player in the game
type Player struct {
	hands       []*Hand
	money       int
	bet         int
	name        string
	isSurrender bool
}

// NewPlayer creates a new Player with the given name and money
func NewPlayer(name string, money int) *Player {
	p := &Player{
		money:       money,
		name:        name,
		isSurrender: false,
	}
	p.ResetHand()
	return p
}

// GetName returns the name of the player
func (p *Player) GetName() string {
	return p.name
}

// AddMoney adds the specified amount to the player's money
func (p *Player) AddMoney(amount int) {
	p.money += amount
}

// GetMoney returns the current amount of money the player has
func (p *Player) GetMoney() int {
	return p.money
}

// ResetHand resets the player's hand and initializes a new hand
func (p *Player) ResetHand() {
	p.isSurrender = false
	p.hands = []*Hand{NewHand()}
}

// GetBet returns the current bet of the player
func (p *Player) GetBet() int {
	return p.bet
}

// SetBet sets the player's bet
func (p *Player) SetBet(bet int) {
	p.bet = bet
}

// GetHands returns the player's hands
func (p *Player) GetHands() []*Hand {
	return p.hands
}

// IsGameOver checks if the player is out of money and thus out of the game
func (p *Player) IsGameOver() bool {
	return p.GetMoney() == 0
}

// CanSplit checks if the player can split their hand
func (p *Player) CanSplit() bool {
	if len(p.hands) != 1 || len(p.hands[0].ToList()) < 2 {
		return false
	}

	firstCard := p.hands[0].ToList()[0]
	secondCard := p.hands[0].ToList()[1]

	return firstCard.GetNumber() == secondCard.GetNumber() && p.money >= p.bet*2
}

// Split splits the player's hand into two hands
func (p *Player) Split(deck *Deck) {
	secondHand := NewHand()
	secondHand.Add(p.hands[0].ToList()[1])
	p.hands = append(p.hands, secondHand)

	// Remove the second card from the first hand
	p.hands[0].cards = p.hands[0].cards[:1]

	// Draw a card for each hand
	p.hands[0].Hit(deck)
	p.hands[1].Hit(deck)
}

// Surrender sets the player's surrender status to true
func (p *Player) Surrender() {
	p.isSurrender = true
}

// IsSurrender checks if the player has surrendered
func (p *Player) IsSurrender() bool {
	return p.isSurrender
}
