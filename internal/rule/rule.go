package rule

import "github.com/panyutsriwirote/GoGo/internal/board"

type ResolutionError struct {
	Reason string
}

func ResolveBoard(board_state *board.BoardState) *ResolutionError {
	return nil
}
