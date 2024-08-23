package cui

import (
	"blackjack/object"
	"bufio"
	"fmt"
	"os"
)

// BlackjackCui represents the main class for playing Blackjack in the console
type BlackjackCui struct {
	game *object.Game
}

// Play starts the Blackjack game
func (b *BlackjackCui) Play() {
	// タイトル表示
	ShowTitle()

	// 標準入力
	scanner := bufio.NewScanner(os.Stdin)

	// ゲーム設定
	gameSettings := GetGameSettings(scanner)
	b.game = object.NewGame(gameSettings)

	// ゲーム終了するまでループ
	for {
		// 掛け金を決める
		SetPlayersBet(scanner, b.game)

		// 最初の手札を配る
		DealInitialHand(b.game)

		// 各プレイヤーのターンを実行する
		DoPlayersTurn(scanner, b.game)

		// 全員バーストしてたら終了する
		if b.game.AllPlayerHandsIsBust() {
			// 結果を表示する
			ShowResultBust(b.game)
			if b.game.AllPlayerIsGameOver() {
				break
			}
			// 続けるか選ぶ
			if GetContinue(scanner, b.game) == 2 {
				break
			}
			continue
		}

		// ディーラーのターンを実行する
		DoDealerTurns(b.game)

		// 結果を表示する
		ShowResult(b.game)
		if b.game.AllPlayerIsGameOver() {
			break
		}

		// 続けるか選ぶ
		if GetContinue(scanner, b.game) == 2 {
			break
		}
	}

	// 最終結果を表示する
	ShowResultEnd(b.game)
	fmt.Println("Good bye.")
}
