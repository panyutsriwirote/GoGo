package board

import "fmt"

const boardSize = 9
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
`

type InvalidCoordError struct {
	coord string
}

func (err *InvalidCoordError) Error() string {
	return fmt.Sprintf("Invalid coordinate: %v", err.coord)
}

func coordTo2DIndex(coord string) (int, int, *InvalidCoordError) {
	if len(coord) != 2 {
		return -1, -1, &InvalidCoordError{coord}
	}

	var x, y int

	if first_component := coord[0]; 'A' <= first_component && first_component <= 'I' {
		y = int(first_component) - 'A'
	} else {
		return -1, -1, &InvalidCoordError{coord}
	}

	if second_component := coord[1]; '1' <= second_component && second_component <= '9' {
		x = '9' - int(second_component)
	} else {
		return -1, -1, &InvalidCoordError{coord}
	}

	return x, y, nil
}

type BoardState struct {
	Prev        *BoardState
	next_player rune
	stones      [boardSize][boardSize]rune
}

func InitBoardState() *BoardState {
	return &BoardState{
		Prev:        nil,
		next_player: 'X',
		stones: [boardSize][boardSize]rune{
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
	coord  string
}

func (err *StonePlacingError) Error() string {
	return fmt.Sprintf("Cannot place stone at %v: %v", err.coord, err.Reason)
}

func (board_state *BoardState) PlaceStone(coord string) (*BoardState, *StonePlacingError) {
	x, y, err := coordTo2DIndex(coord)
	if err != nil {
		return board_state, &StonePlacingError{
			Reason: "Invalid coordinate",
			coord:  coord,
		}
	}
	if board_state.stones[x][y] != ' ' {
		return board_state, &StonePlacingError{
			Reason: "Space already taken",
			coord:  coord,
		}
	}
	new_stones := board_state.stones
	new_stones[x][y] = board_state.next_player
	var new_next_player rune
	if board_state.next_player == 'X' {
		new_next_player = 'O'
	} else {
		new_next_player = 'X'
	}
	new_state := BoardState{
		Prev:        board_state,
		next_player: new_next_player,
		stones:      new_stones,
	}
	return &new_state, nil
}

func (board_state *BoardState) Display() {
	fmt.Printf(
		boardDisplayTemplate,
		board_state.stones[0][0],
		board_state.stones[0][1],
		board_state.stones[0][2],
		board_state.stones[0][3],
		board_state.stones[0][4],
		board_state.stones[0][5],
		board_state.stones[0][6],
		board_state.stones[0][7],
		board_state.stones[0][8],
		board_state.stones[1][0],
		board_state.stones[1][1],
		board_state.stones[1][2],
		board_state.stones[1][3],
		board_state.stones[1][4],
		board_state.stones[1][5],
		board_state.stones[1][6],
		board_state.stones[1][7],
		board_state.stones[1][8],
		board_state.stones[2][0],
		board_state.stones[2][1],
		board_state.stones[2][2],
		board_state.stones[2][3],
		board_state.stones[2][4],
		board_state.stones[2][5],
		board_state.stones[2][6],
		board_state.stones[2][7],
		board_state.stones[2][8],
		board_state.stones[3][0],
		board_state.stones[3][1],
		board_state.stones[3][2],
		board_state.stones[3][3],
		board_state.stones[3][4],
		board_state.stones[3][5],
		board_state.stones[3][6],
		board_state.stones[3][7],
		board_state.stones[3][8],
		board_state.stones[4][0],
		board_state.stones[4][1],
		board_state.stones[4][2],
		board_state.stones[4][3],
		board_state.stones[4][4],
		board_state.stones[4][5],
		board_state.stones[4][6],
		board_state.stones[4][7],
		board_state.stones[4][8],
		board_state.stones[5][0],
		board_state.stones[5][1],
		board_state.stones[5][2],
		board_state.stones[5][3],
		board_state.stones[5][4],
		board_state.stones[5][5],
		board_state.stones[5][6],
		board_state.stones[5][7],
		board_state.stones[5][8],
		board_state.stones[6][0],
		board_state.stones[6][1],
		board_state.stones[6][2],
		board_state.stones[6][3],
		board_state.stones[6][4],
		board_state.stones[6][5],
		board_state.stones[6][6],
		board_state.stones[6][7],
		board_state.stones[6][8],
		board_state.stones[7][0],
		board_state.stones[7][1],
		board_state.stones[7][2],
		board_state.stones[7][3],
		board_state.stones[7][4],
		board_state.stones[7][5],
		board_state.stones[7][6],
		board_state.stones[7][7],
		board_state.stones[7][8],
		board_state.stones[8][0],
		board_state.stones[8][1],
		board_state.stones[8][2],
		board_state.stones[8][3],
		board_state.stones[8][4],
		board_state.stones[8][5],
		board_state.stones[8][6],
		board_state.stones[8][7],
		board_state.stones[8][8],
	)
}
