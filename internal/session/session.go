package session

import (
	"fmt"
	"strings"

	"github.com/panyutsriwirote/GoGo/internal/board"
	"github.com/panyutsriwirote/GoGo/internal/rule"
)

type GameSession struct {
	Board *board.BoardState
}

func NewGameSession() *GameSession {
	return &GameSession{
		Board: board.InitBoardState(),
	}
}

type GameEndSignal struct {
	Reason     string
	FatalError bool
}

func (end_signal *GameEndSignal) String() string {
	format := "GAME ENDED due to: "
	if end_signal.FatalError {
		format += "FATAL ERROR: "
	}
	format += "%v"
	return fmt.Sprintf(format, end_signal.Reason)
}

func (session *GameSession) PlayTurn() *GameEndSignal {
	fmt.Printf("%c's move: ", session.Board.NextPlayer)
	var move string
	var read_err error
	_, read_err = fmt.Scanln(&move)
	if read_err != nil {
		fmt.Print("\n")
		return &GameEndSignal{
			Reason:     "Failed to read user input",
			FatalError: true,
		}
	}

	switch move = strings.ToUpper(move); move {
	case "U":
		if session.Board.Prev == nil {
			fmt.Println("Cannot undo further!")
		} else {
			session.Board = session.Board.Prev
		}
	case "P":
		session.Board = session.Board.Pass()
		if session.Board.Prev.LastMove == "P" {
			return &GameEndSignal{
				Reason:     "2 consecutive passes",
				FatalError: false,
			}
		}
	default:
		var stone_err *board.StonePlacingError
		session.Board, stone_err = session.Board.PlaceStone(move)
		if stone_err == nil {
			rule.ResolveBoard(session.Board)
		} else {
			fmt.Printf("%v. Please try again.\n", stone_err.Reason)
		}
	}
	return nil
}

func (session *GameSession) CountScores() (int, int) {
	return session.Board.XPrisoner, session.Board.OPrisoner
}
