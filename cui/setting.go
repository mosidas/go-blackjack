package cui

import (
	"blackjack/object"
	"bufio"
	"fmt"
	"strconv"
)

// GetGameSettings prompts the user to input game settings
func GetGameSettings(scanner *bufio.Scanner) *object.GameSettings {
	// プレイヤー人数設定
	playerCount := setPlayersNumber(scanner)
	// デッキ数設定
	deckCount := setDeckNumber(scanner)
	// ソフト17ルール設定
	soft17 := setHitSoft17(scanner)
	// 所持金設定
	defaultMoney := setDefaultMoney(scanner)

	return object.NewGameSettings(playerCount, deckCount, soft17, defaultMoney)
}

func setPlayersNumber(scanner *bufio.Scanner) int {
	for {
		fmt.Print("プレイヤー人数:1～8 →   ")
		scanner.Scan()
		input, err := strconv.Atoi(scanner.Text())
		if err == nil && input >= 1 && input <= 8 {
			return input
		}
		fmt.Println("1～8の整数値を入力してください")
	}
}

func setDeckNumber(scanner *bufio.Scanner) int {
	for {
		fmt.Print("デッキ数:1～8 →   ")
		scanner.Scan()
		input, err := strconv.Atoi(scanner.Text())
		if err == nil && input >= 1 && input <= 8 {
			return input
		}
		fmt.Println("1～8の整数値を入力してください")
	}
}

func setHitSoft17(scanner *bufio.Scanner) bool {
	for {
		fmt.Print("ソフト17(1:ヒットする  2:スタンドする):→   ")
		scanner.Scan()
		input, err := strconv.Atoi(scanner.Text())
		if err == nil {
			if input == 1 {
				return true
			} else if input == 2 {
				return false
			}
		}
		fmt.Println("1か2を入力してください")
	}
}

func setDefaultMoney(scanner *bufio.Scanner) int {
	for {
		fmt.Print("最初の所持金(10～100000):→   ")
		scanner.Scan()
		input, err := strconv.Atoi(scanner.Text())
		if err == nil && input >= 10 && input <= 100000 {
			return input
		}
		fmt.Println("10～100000の整数値を入力してください")
	}
}
