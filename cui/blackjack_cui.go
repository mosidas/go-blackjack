package cui

import (
	"fmt"
)

func Run() {
	for {
		showTitle()
		inputGameSetting()
		inputPlayerBet()
		showInitialHands()
		doPlayerTurn()
		doDealerTurn()
		showResult()
	}
}

func showTitle() {
	fmt.Println("Blackjack")
}

func inputGameSetting() {
	fmt.Println("inputGameSetting")
}

func inputPlayerBet() {
	fmt.Println("inputPlayerBet")
}

func showInitialHands() {
	fmt.Println("showInitialHands")
}

func doPlayerTurn() {
	fmt.Println("doPlayerTurn")
}

func doDealerTurn() {
	fmt.Println("doDealerTurn")
}

func showResult() {
	fmt.Println("showResult")
}
