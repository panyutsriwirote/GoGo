package rule

import (
	"github.com/panyutsriwirote/GoGo/internal/board"
	"github.com/panyutsriwirote/GoGo/internal/coord"
)

type ResolutionError struct {
	Reason string
}

func ResolveBoard(board_state *board.BoardState) *ResolutionError {
	next_player := board_state.NextPlayer
	prev_player := board.SwitchPlayer(next_player)

	// Remove captured groups.
	for x := range board_state.Stones {
		for y := range board_state.Stones[x] {
			pos := coord.New(x, y)
			if board_state.GetOwner(pos) == next_player {
				group := board_state.GetGroup(pos)
				group_liberty := board_state.GetGroupLiberty(group)
				if group_liberty == 0 {
					board_state.RemoveGroup(group)
					board_state.AddPrisoner(prev_player, len(group))
				}
			}
		}
	}

	// Check suicide.
	for x := range board_state.Stones {
		for y := range board_state.Stones[x] {
			pos := coord.New(x, y)
			if board_state.GetOwner(pos) == prev_player {
				group := board_state.GetGroup(pos)
				group_liberty := board_state.GetGroupLiberty(group)
				if group_liberty == 0 {
					return &ResolutionError{"Suicide"}
				}
			}
		}
	}

	// Check Ko.
	if board_state.Prev.Prev != nil &&
		board_state.Stones == board_state.Prev.Prev.Stones {
		return &ResolutionError{"KO"}
	}

	return nil
}

type GameScore struct {
	X, O int
}

func CountScores(board_state *board.BoardState) GameScore {
	return GameScore{
		X: board_state.XPrisoner,
		O: board_state.OPrisoner,
	}
}
