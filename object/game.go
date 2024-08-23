package object

import (
	"fmt"
)

// Results represents the outcome of a game
type Results int

const (
	Win Results = iota
	Lose
	Draw
	Surrender
)

// Game represents the game management
type Game struct {
	deck     *Deck
	players  []*Player
	dealer   *Player
	settings *GameSettings
}

// NewGame creates a new Game with the given settings
func NewGame(settings *GameSettings) *Game {
	game := &Game{
		deck:     NewDeck(),
		players:  []*Player{},
		dealer:   NewPlayer("ディーラー", 0),
		settings: settings,
	}
	game.initPlayers(game.settings.GetPlayerCount(), game.settings.GetDefaultMoney())
	game.deck.SetDeckNumber(game.settings.GetDeckCount())
	return game
}

// initPlayers initializes the players with the given count and money
func (g *Game) initPlayers(count, money int) {
	for i := 0; i < count; i++ {
		g.players = append(g.players, NewPlayer("プレイヤー"+fmt.Sprintf("%d", i+1), money))
	}
}

// GetDealer returns the dealer of the game
func (g *Game) GetDealer() *Player {
	return g.dealer
}

// GetActivePlayers returns a list of active players (not game over)
func (g *Game) GetActivePlayers() []*Player {
	activePlayers := []*Player{}
	for _, player := range g.players {
		if !player.IsGameOver() {
			activePlayers = append(activePlayers, player)
		}
	}
	return activePlayers
}

// GetAllPlayers returns a list of all players in the game
func (g *Game) GetAllPlayers() []*Player {
	return g.players
}

// GetDeck returns the deck used in the game
func (g *Game) GetDeck() *Deck {
	return g.deck
}

// IsSoft17Hit returns whether the soft 17 rule is applied
func (g *Game) IsSoft17Hit() bool {
	return g.settings.GetSoft17()
}

// DefaultMoney returns the default money for the players
func (g *Game) DefaultMoney() int {
	return g.settings.GetDefaultMoney()
}

// DealInitialHand deals the initial two cards to each player and the dealer
func (g *Game) DealInitialHand() {
	for _, player := range g.GetActivePlayers() {
		player.ResetHand()
		player.GetHands()[0].Hit(g.GetDeck())
		player.GetHands()[0].Hit(g.GetDeck())
	}
	g.dealer.ResetHand()
	g.dealer.GetHands()[0].Hit(g.GetDeck())
	g.dealer.GetHands()[0].Hit(g.GetDeck())
}

// GetResult determines the result of the game for the given player and hand
func (g *Game) GetResult(player *Player, hand *Hand) Results {
	dealerHand := g.dealer.GetHands()[0]
	if player.IsSurrender() {
		return Surrender
	}
	if hand.IsBust() {
		return Lose
	} else if dealerHand.IsBust() {
		return Win
	} else {
		if hand.IsNaturalBlackjack() && !dealerHand.IsNaturalBlackjack() {
			return Win
		} else if hand.IsNaturalBlackjack() && dealerHand.IsNaturalBlackjack() {
			return Draw
		} else if !hand.IsNaturalBlackjack() && dealerHand.IsNaturalBlackjack() {
			return Lose
		} else if hand.GetScore() > dealerHand.GetScore() {
			return Win
		} else if hand.GetScore() < dealerHand.GetScore() {
			return Lose
		} else {
			return Draw
		}
	}
}

// CalcReturn calculates the return for the player based on the game result
func (g *Game) CalcReturn(player *Player, hand *Hand) {
	result := g.GetResult(player, hand)
	switch result {
	case Win:
		player.AddMoney(player.GetBet())
	case Lose:
		player.AddMoney(-player.GetBet())
	case Surrender:
		player.AddMoney(-player.GetBet() / 2)
	}
}

// AllPlayerHandsIsBust checks if all players' hands are bust
func (g *Game) AllPlayerHandsIsBust() bool {
	for _, player := range g.GetActivePlayers() {
		for _, hand := range player.GetHands() {
			if !hand.IsBust() {
				return false
			}
		}
	}
	return true
}

// AllPlayerIsGameOver checks if all players are out of the game
func (g *Game) AllPlayerIsGameOver() bool {
	return len(g.GetActivePlayers()) == 0
}
