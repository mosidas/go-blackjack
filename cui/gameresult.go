package cui

import (
	"blackjack/object"
	"fmt"
	"sort"
	"time"
)

// ShowResult displays the game results
func ShowResult(gameManager *object.Game) {
	fmt.Println("------------結果------------")
	time.Sleep(1 * time.Second)
	fmt.Print("ディーラーの点数：")
	fmt.Println(gameManager.GetDealer().GetHands()[0].GetScore())

	for _, player := range gameManager.GetActivePlayers() {
		time.Sleep(1 * time.Second)
		for i, hand := range player.GetHands() {
			if len(player.GetHands()) >= 2 {
				fmt.Printf("%sの手札%dの点数：", player.GetName(), i+1)
			} else {
				fmt.Printf("%sの点数：", player.GetName())
			}
			fmt.Println(hand.GetScore())
			gameManager.CalcReturn(player, hand)
			time.Sleep(1 * time.Second)
			switch gameManager.GetResult(player, hand) {
			case object.Win:
				fmt.Println("勝ちました！！！おめでとう！！！")
				fmt.Printf("お金：%d( +%d )\n", player.GetMoney(), player.GetBet())
			case object.Lose:
				fmt.Println("負けました...")
				fmt.Printf("お金：%d( -%d )\n", player.GetMoney(), player.GetBet())
				if player.IsGameOver() {
					fmt.Printf("%s、 ゲームオーバー！\n", player.GetName())
				}
			default:
				fmt.Println("引き分けです！")
				fmt.Printf("お金：%d( +-0 )\n", player.GetMoney())
			}
		}
	}

	fmt.Println("----------------------------")
	time.Sleep(1 * time.Second)
}

// ShowResultBust displays the game results when all players bust
func ShowResultBust(gameManager *object.Game) {
	fmt.Println("------------結果------------")
	for _, player := range gameManager.GetActivePlayers() {
		for i, hand := range player.GetHands() {
			if len(player.GetHands()) >= 2 {
				fmt.Printf("%sの手札%dの点数：", player.GetName(), i+1)
			} else {
				fmt.Printf("%sの点数：", player.GetName())
			}
			time.Sleep(1 * time.Second)
			fmt.Printf("%sの点数：", player.GetName())
			fmt.Println(hand.GetScore())
			time.Sleep(1 * time.Second)
			gameManager.CalcReturn(player, hand)
			fmt.Printf("お金：%d( -%d )\n", player.GetMoney(), player.GetBet())
			if player.IsGameOver() {
				fmt.Printf("%s、 ゲームオーバー！\n", player.GetName())
			}
		}
	}
	fmt.Println("----------------------------")
	time.Sleep(1 * time.Second)
}

// ShowResultEnd displays the final results at the end of the game
func ShowResultEnd(gameManager *object.Game) {
	fmt.Print("最終結果")
	time.Sleep(500 * time.Millisecond)
	fmt.Print(".")
	time.Sleep(500 * time.Millisecond)
	fmt.Print(".")
	time.Sleep(500 * time.Millisecond)
	fmt.Println(".")

	players := gameManager.GetAllPlayers()
	sort.Slice(players, func(i, j int) bool {
		return players[i].GetMoney() > players[j].GetMoney()
	})

	for count, player := range players {
		fmt.Printf("%d:%s   所持金：%d(", count+1, player.GetName(), player.GetMoney())
		delta := player.GetMoney() - gameManager.DefaultMoney()
		if delta > 0 {
			fmt.Print("+")
		} else if delta == 0 {
			fmt.Print("+-")
		}
		fmt.Printf("%d)\n", delta)
		time.Sleep(1 * time.Second)
	}
}
