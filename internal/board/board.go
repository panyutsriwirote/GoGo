package board

import (
	"fmt"
	"strings"
)

const BoardSize = 9
const boardDisplayTemplate = `
    A   B   C   D   E   F   G   H   I
  +---+---+---+---+---+---+---+---+---+
9 | %c | %c | %c | %c | %c | %c | %c | %c | %c |
  +---+---+---+---+---+---+---+---+---+
8 | %c | %c | %c | %c | %c | %c | %c | %c | %c |
  +---+---+---+---+---+---+---+---+---+
7 | %c | %c | %c | %c | %c | %c | %c | %c | %c |
  +---+---+---+---+---+---+---+---+---+
6 | %c | %c | %c | %c | %c | %c | %c | %c | %c |
  +---+---+---+---+---+---+---+---+---+
5 | %c | %c | %c | %c | %c | %c | %c | %c | %c |
  +---+---+---+---+---+---+---+---+---+
4 | %c | %c | %c | %c | %c | %c | %c | %c | %c |
  +---+---+---+---+---+---+---+---+---+
3 | %c | %c | %c | %c | %c | %c | %c | %c | %c |
  +---+---+---+---+---+---+---+---+---+
2 | %c | %c | %c | %c | %c | %c | %c | %c | %c |
  +---+---+---+---+---+---+---+---+---+
1 | %c | %c | %c | %c | %c | %c | %c | %c | %c |
  +---+---+---+---+---+---+---+---+---+
             X's prisoner: %v
             O's prisoner: %v
`

type InvalidCoordError struct {
	coord_string string
}

func (err *InvalidCoordError) Error() string {
	return fmt.Sprintf("Invalid coordinate: %v", err.coord_string)
}

type Coord struct {
	X   int
	Y   int
	Str string
}

func (coord *Coord) String() string {
	return coord.Str
}

func stringToCoord(coord_string string) (*Coord, *InvalidCoordError) {
	if len(coord_string) != 2 {
		return nil, &InvalidCoordError{coord_string}
	}

	coord_string = strings.ToUpper(coord_string)

	var x, y int

	if first_component := coord_string[0]; 'A' <= first_component && first_component <= 'I' {
		y = int(first_component) - 'A'
	} else {
		return nil, &InvalidCoordError{coord_string}
	}

	if second_component := coord_string[1]; '1' <= second_component && second_component <= '9' {
		x = '9' - int(second_component)
	} else {
		return nil, &InvalidCoordError{coord_string}
	}

	return &Coord{x, y, coord_string}, nil
}

type BoardState struct {
	Prev       *BoardState
	LastMove   string
	NextPlayer rune
	XPrisoner  int
	OPrisoner  int
	Stones     [BoardSize][BoardSize]rune
}

func InitBoardState() *BoardState {
	return &BoardState{
		Prev:       nil,
		LastMove:   "",
		NextPlayer: 'X',
		XPrisoner:  0,
		OPrisoner:  0,
		Stones: [BoardSize][BoardSize]rune{
			{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
			{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
			{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
			{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
			{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
			{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
			{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
			{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
			{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
		},
	}
}

type StonePlacingError struct {
	Reason string
	coord  *Coord
}

func (err *StonePlacingError) Error() string {
	return fmt.Sprintf("Cannot place stone at %v: %v", err.coord, err.Reason)
}

func (board_state *BoardState) PlaceStone(coord_string string) (*BoardState, *StonePlacingError) {
	coord, err := stringToCoord(coord_string)
	if err != nil {
		return board_state, &StonePlacingError{
			Reason: "Invalid coordinate",
			coord:  coord,
		}
	}
	if board_state.Stones[coord.X][coord.Y] != ' ' {
		return board_state, &StonePlacingError{
			Reason: "Space already taken",
			coord:  coord,
		}
	}
	new_stones := board_state.Stones
	new_stones[coord.X][coord.Y] = board_state.NextPlayer
	new_state := BoardState{
		Prev:       board_state,
		LastMove:   coord.Str,
		NextPlayer: switchPlayer(board_state.NextPlayer),
		XPrisoner:  board_state.XPrisoner,
		OPrisoner:  board_state.OPrisoner,
		Stones:     new_stones,
	}
	return &new_state, nil
}

func (board_state *BoardState) Pass() *BoardState {
	return &BoardState{
		Prev:       board_state,
		LastMove:   "P",
		NextPlayer: switchPlayer(board_state.NextPlayer),
		XPrisoner:  board_state.XPrisoner,
		OPrisoner:  board_state.OPrisoner,
		Stones:     board_state.Stones,
	}
}

func switchPlayer(player rune) rune {
	if player == 'X' {
		return 'O'
	} else {
		return 'X'
	}
}

func (board_state *BoardState) Display() {
	fmt.Printf(
		boardDisplayTemplate,
		board_state.Stones[0][0],
		board_state.Stones[0][1],
		board_state.Stones[0][2],
		board_state.Stones[0][3],
		board_state.Stones[0][4],
		board_state.Stones[0][5],
		board_state.Stones[0][6],
		board_state.Stones[0][7],
		board_state.Stones[0][8],
		board_state.Stones[1][0],
		board_state.Stones[1][1],
		board_state.Stones[1][2],
		board_state.Stones[1][3],
		board_state.Stones[1][4],
		board_state.Stones[1][5],
		board_state.Stones[1][6],
		board_state.Stones[1][7],
		board_state.Stones[1][8],
		board_state.Stones[2][0],
		board_state.Stones[2][1],
		board_state.Stones[2][2],
		board_state.Stones[2][3],
		board_state.Stones[2][4],
		board_state.Stones[2][5],
		board_state.Stones[2][6],
		board_state.Stones[2][7],
		board_state.Stones[2][8],
		board_state.Stones[3][0],
		board_state.Stones[3][1],
		board_state.Stones[3][2],
		board_state.Stones[3][3],
		board_state.Stones[3][4],
		board_state.Stones[3][5],
		board_state.Stones[3][6],
		board_state.Stones[3][7],
		board_state.Stones[3][8],
		board_state.Stones[4][0],
		board_state.Stones[4][1],
		board_state.Stones[4][2],
		board_state.Stones[4][3],
		board_state.Stones[4][4],
		board_state.Stones[4][5],
		board_state.Stones[4][6],
		board_state.Stones[4][7],
		board_state.Stones[4][8],
		board_state.Stones[5][0],
		board_state.Stones[5][1],
		board_state.Stones[5][2],
		board_state.Stones[5][3],
		board_state.Stones[5][4],
		board_state.Stones[5][5],
		board_state.Stones[5][6],
		board_state.Stones[5][7],
		board_state.Stones[5][8],
		board_state.Stones[6][0],
		board_state.Stones[6][1],
		board_state.Stones[6][2],
		board_state.Stones[6][3],
		board_state.Stones[6][4],
		board_state.Stones[6][5],
		board_state.Stones[6][6],
		board_state.Stones[6][7],
		board_state.Stones[6][8],
		board_state.Stones[7][0],
		board_state.Stones[7][1],
		board_state.Stones[7][2],
		board_state.Stones[7][3],
		board_state.Stones[7][4],
		board_state.Stones[7][5],
		board_state.Stones[7][6],
		board_state.Stones[7][7],
		board_state.Stones[7][8],
		board_state.Stones[8][0],
		board_state.Stones[8][1],
		board_state.Stones[8][2],
		board_state.Stones[8][3],
		board_state.Stones[8][4],
		board_state.Stones[8][5],
		board_state.Stones[8][6],
		board_state.Stones[8][7],
		board_state.Stones[8][8],
		board_state.XPrisoner,
		board_state.OPrisoner,
	)
}
