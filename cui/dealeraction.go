package cui

import (
	"blackjack/object"
	"fmt"
	"time"
)

// DealInitialHand deals the initial hand for all players and the dealer
func DealInitialHand(gameManager *object.Game) {
	// 山札を作成する。
	gameManager.GetDeck().Create()
	// 手札を配る。
	gameManager.DealInitialHand()

	// 各プレイヤーの手札を表示する。
	for _, player := range gameManager.GetActivePlayers() {
		fmt.Printf("%sの手札：", player.GetName())
		for _, card := range player.GetHands()[0].ToList() {
			fmt.Print(card.GetText() + " ")
		}
		fmt.Println("")
	}

	// ディーラーの手札を表示する。
	fmt.Print("ディーラーの手札：")
	dealerCards := gameManager.GetDealer().GetHands()[0].ToList()
	fmt.Printf("%s ??\n", dealerCards[0].GetText())
}

// DoDealerTurns displays and executes the dealer's turn
func DoDealerTurns(gameManager *object.Game) {
	fmt.Print("ディーラーのターン.")
	time.Sleep(500 * time.Millisecond)
	fmt.Print(".")
	time.Sleep(500 * time.Millisecond)
	fmt.Print(".")
	time.Sleep(500 * time.Millisecond)
	fmt.Println(".")

	hand := gameManager.GetDealer().GetHands()[0]
	for !hand.IsBust() &&
		(hand.GetScore() < 17 || (hand.IsSoft17() && gameManager.IsSoft17Hit())) &&
		gameManager.GetDeck().Size() > 0 {

		hand.Hit(gameManager.GetDeck())
	}

	// ディーラーの手札を表示する。
	fmt.Print("ディーラーの手札：")
	for _, card := range hand.ToList() {
		fmt.Print(card.GetText() + " ")
	}
	fmt.Println("")
	time.Sleep(1 * time.Second)
}
