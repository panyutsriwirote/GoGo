package main

import (
	"fmt"

	"github.com/panyutsriwirote/GoGo/internal/session"
)

func main() {
	game_session := session.New()
	game_session.Board.Display()
	for {
		end_signal := game_session.PlayTurn()
		if end_signal != nil {
			fmt.Println(end_signal)
			if end_signal.FatalError {
				return
			} else {
				break
			}
		}
		game_session.Board.Display()
	}
	score := game_session.CountScores()
	fmt.Println("Scores:")
	fmt.Printf("\tX: %v\n", score.X)
	fmt.Printf("\tO: %v\n", score.O)
}
