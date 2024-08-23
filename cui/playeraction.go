package cui

import (
	"blackjack/object"
	"bufio"
	"fmt"
	"strconv"
	"time"
)

// SetPlayersBet sets the bet amount for each active player
func SetPlayersBet(scanner *bufio.Scanner, gameManager *object.Game) {
	for _, player := range gameManager.GetActivePlayers() {
		for {
			fmt.Printf("%sのベット(現在の所持金：%d)：", player.GetName(), player.GetMoney())
			scanner.Scan()
			input, err := strconv.Atoi(scanner.Text())
			if err == nil && input >= 1 && input <= player.GetMoney() {
				player.SetBet(input)
				break
			}
			fmt.Println("1以上、所持金以下の整数値を入力してください")
		}
	}
}

// DoPlayersTurn executes each player's turn
func DoPlayersTurn(scanner *bufio.Scanner, gameManager *object.Game) {
	for _, player := range gameManager.GetActivePlayers() {
		DoPlayerTurns(scanner, gameManager, player)
	}
}

// DoPlayerTurns executes the turn for a single player
func DoPlayerTurns(scanner *bufio.Scanner, gameManager *object.Game, player *object.Player) {
	stand := false

	fmt.Printf("%sのターン！\n", player.GetName())
	fmt.Printf("%sの手札：", player.GetName())
	hand := player.GetHands()[0]
	for _, card := range hand.ToList() {
		fmt.Print(card.GetText() + " ")
	}
	fmt.Println("")

	// バーストするか、スタンドするか、山札がなくなるまでヒットかスタンドかを選ぶ
	for !hand.IsBust() && !stand && gameManager.GetDeck().Size() > 0 {
		input := getPlayerAction(scanner, gameManager, player)
		switch input {
		case 1:
			// ヒット
			hand.Hit(gameManager.GetDeck())
		case 2:
			// スタンド
			stand = true
		case 3:
			// サレンダー
			player.Surrender()
			return
		case 4:
			// スプリット
			DoSplitedPlayerTurns(scanner, gameManager, player)
			return
		}

		fmt.Printf("%sの手札：", player.GetName())
		for _, card := range hand.ToList() {
			fmt.Print(card.GetText() + " ")
		}
		fmt.Println("")
	}

	// バーストしてたらゲームオーバー。
	if hand.IsBust() {
		showResultBust(gameManager, player, hand)
	}
}

// DoSplitedPlayerTurns executes the turns for a player who has split their hand
func DoSplitedPlayerTurns(scanner *bufio.Scanner, gameManager *object.Game, player *object.Player) {
	player.Split(gameManager.GetDeck())
	fmt.Printf("%sはスプリットした！\n", player.GetName())
	fmt.Printf("%sの手札：", player.GetName())
	for _, hand := range player.GetHands() {
		for _, card := range hand.ToList() {
			fmt.Print(card.GetText() + " ")
		}
		fmt.Println("")
	}
	i := 1
	for _, hand := range player.GetHands() {
		stand := false
		fmt.Printf("%sの手札%dのターン!\n", player.GetName(), i)
		fmt.Printf("%sの手札%d：", player.GetName(), i)
		for _, card := range hand.ToList() {
			fmt.Print(card.GetText() + " ")
		}
		fmt.Println("")

		for !hand.IsBust() && !stand && gameManager.GetDeck().Size() > 0 {
			input := getPlayerAction(scanner, gameManager, player)
			switch input {
			case 1:
				// ヒット
				hand.Hit(gameManager.GetDeck())
			case 2:
				// スタンド
				stand = true
			}

			fmt.Printf("%sの手札%d：", player.GetName(), i)
			for _, card := range hand.ToList() {
				fmt.Print(card.GetText() + " ")
			}
			fmt.Println("")
		}

		if hand.IsBust() {
			showResultBust(gameManager, player, hand)
		}
		i++
	}
}

// getPlayerAction gets the player's action: Hit, Stand, Surrender, or Split
func getPlayerAction(scanner *bufio.Scanner, gameManager *object.Game, player *object.Player) int {
	for {
		fmt.Print("どうする? ヒット:1 スタンド:2")
		if canSurrender(gameManager, player) {
			fmt.Print(" サレンダー:3")
		}
		if player.CanSplit() {
			fmt.Print(" スプリット:4")
		}
		if canDoubledown(gameManager, player) {
			fmt.Print(" ダブルダウン:5")
		}
		fmt.Print(" →   ")

		scanner.Scan()
		input, err := strconv.Atoi(scanner.Text())
		if err == nil && (input == 1 || input == 2 || (player.CanSplit() && input == 4)) {
			return input
		}
		fmt.Println("選択可能な数値を入力してください")
	}
}

func canDoubledown(gameManager *object.Game, player *object.Player) bool {
	// TODO: Implement this logic
	return false
}

func canSurrender(gameManager *object.Game, player *object.Player) bool {
	// TODO: Implement this logic
	return false
}

// showResultBust displays the result when a player busts
func showResultBust(gameManager *object.Game, player *object.Player, hand *object.Hand) {
	fmt.Printf("%sの点数：%d\n", player.GetName(), hand.GetScore())
	time.Sleep(1 * time.Second)
	fmt.Println("バースト！")
	time.Sleep(1 * time.Second)
}
