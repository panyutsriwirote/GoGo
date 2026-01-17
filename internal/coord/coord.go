package coord

import (
	"fmt"
	"strings"

	"github.com/panyutsriwirote/GoGo/internal/constants"
)

type Coord struct {
	X, Y int
}

var alreadyCreated [constants.BOARD_SIZE][constants.BOARD_SIZE]*Coord

func New(x, y int) *Coord {
	if x < 0 ||
		y < 0 ||
		x >= constants.BOARD_SIZE ||
		y >= constants.BOARD_SIZE {
		return nil
	}
	existing_coord := alreadyCreated[x][y]
	if existing_coord != nil {
		return existing_coord
	}
	new_coord := &Coord{x, y}
	alreadyCreated[x][y] = new_coord
	return new_coord
}

func (c *Coord) String() string {
	return fmt.Sprintf(
		"%c%v",
		constants.HORIZONTAL_COORD_START+c.Y,
		constants.BOARD_SIZE-c.X,
	)
}

func FromString(coord_string string) *Coord {
	if len(coord_string) != 2 {
		return nil
	}

	coord_string = strings.ToUpper(coord_string)

	var x, y int

	first_component := coord_string[0]
	if constants.HORIZONTAL_COORD_START <= first_component &&
		first_component <= constants.HORIZONTAL_COORD_STOP {
		y = int(first_component) - constants.HORIZONTAL_COORD_START
	} else {
		return nil
	}

	second_component := coord_string[1]
	if constants.VERTICAL_COORD_START <= second_component &&
		second_component <= constants.VERTICAL_COORD_STOP {
		x = constants.VERTICAL_COORD_STOP - int(second_component)
	} else {
		return nil
	}

	return New(x, y)
}

func (c *Coord) North() *Coord { return New(c.X-1, c.Y) }

func (c *Coord) South() *Coord { return New(c.X+1, c.Y) }

func (c *Coord) East() *Coord { return New(c.X, c.Y+1) }

func (c *Coord) West() *Coord { return New(c.X, c.Y-1) }
