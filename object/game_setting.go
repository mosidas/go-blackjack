package object

// GameSettings represents the settings for the game
type GameSettings struct {
	playerCount  int
	deckCount    int
	soft17       bool
	defaultMoney int
}

// NewGameSettings creates a new GameSettings with the given parameters
func NewGameSettings(pc int, dc int, s17 bool, dm int) *GameSettings {
	return &GameSettings{
		playerCount:  pc,
		deckCount:    dc,
		soft17:       s17,
		defaultMoney: dm,
	}
}

// GetPlayerCount returns the number of players
func (gs *GameSettings) GetPlayerCount() int {
	return gs.playerCount
}

// GetDeckCount returns the number of decks
func (gs *GameSettings) GetDeckCount() int {
	return gs.deckCount
}

// GetSoft17 returns whether soft 17 is considered in the game
func (gs *GameSettings) GetSoft17() bool {
	return gs.soft17
}

// GetDefaultMoney returns the default money each player starts with
func (gs *GameSettings) GetDefaultMoney() int {
	return gs.defaultMoney
}
