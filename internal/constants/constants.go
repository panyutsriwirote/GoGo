package constants

const BOARD_SIZE = 9
const BOARD_DISPLAY_TEMPLATE = `
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
const HORIZONTAL_COORD_START = 'A'
const HORIZONTAL_COORD_STOP = 'I'
const VERTICAL_COORD_START = '1'
const VERTICAL_COORD_STOP = '9'
