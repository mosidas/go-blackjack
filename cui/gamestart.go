package cui

import (
	"blackjack/object"
	"bufio"
	"fmt"
	"strconv"
)

// GetContinue prompts the user to decide whether to continue or end the game
func GetContinue(scanner *bufio.Scanner, gameManager *object.Game) int {
	if len(gameManager.GetActivePlayers()) == 0 {
		return 2
	}

	fmt.Println("続けますか?")
	for {
		fmt.Print("続ける:1 終わる:2 →   ")
		scanner.Scan()
		input, err := strconv.Atoi(scanner.Text())
		if err == nil && (input == 1 || input == 2) {
			return input
		}
		fmt.Println("1か2を入力してください")
	}
}
