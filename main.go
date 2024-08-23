package main

import (
	"blackjack/cui"
)

// main function serves as the entry point for the Blackjack game
func main() {
	bjCui := &cui.BlackjackCui{}
	bjCui.Play()
}
