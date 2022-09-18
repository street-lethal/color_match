package src

import "fmt"

type Color string

const (
	Red    Color = "R"
	Yellow Color = "Y"
	Green  Color = "G"
	Blue   Color = "B"
	Pink   Color = "P"
)

type Tile interface {
	Copy() Tile
	Display() []string
	OverlapsWith(other Tile) bool
	Rotate90()
	MatchesRight(other Tile) bool
	MatchesUp(other Tile) bool
	MatchesLeft(other Tile) bool
	MatchesDown(other Tile) bool
	AdjustToRight(other Tile) bool
	AdjustToUp(other Tile) bool
	AdjustToLeft(other Tile) bool
	AdjustToDown(other Tile) bool
}
type tile struct {
	Right, Up, Left, Down Color
}

func NewTile(right, up, left, down Color) Tile {
	return &tile{right, up, left, down}
}

func (t *tile) Copy() Tile {
	return NewTile(t.Right, t.Up, t.Left, t.Down)
}

func (t *tile) Display() []string {
	return []string{
		fmt.Sprintf("   %v   ", t.Up),
		fmt.Sprintf(" %v   %v ", t.Left, t.Right),
		fmt.Sprintf("   %v   ", t.Down),
	}
}

func (t *tile) OverlapsWith(other Tile) bool {
	otherTile, _ := other.Copy().(*tile)

	return t.Right == otherTile.Right &&
		t.Down == otherTile.Down &&
		t.Left == otherTile.Left &&
		t.Up == otherTile.Up
}

func (t *tile) Rotate90() {
	t.Right, t.Up, t.Left, t.Down = t.Down, t.Right, t.Up, t.Left
}

func (t *tile) Rotate180() {
	t.Right, t.Up, t.Left, t.Down = t.Left, t.Down, t.Right, t.Up
}

func (t *tile) Rotate270() {
	t.Right, t.Up, t.Left, t.Down = t.Up, t.Left, t.Down, t.Right
}

func (t *tile) MatchesRight(other Tile) bool {
	otherTile, _ := other.(*tile)

	return t.Right == otherTile.Left
}

func (t *tile) MatchesLeft(other Tile) bool {
	otherTile, _ := other.(*tile)

	return t.Left == otherTile.Right
}

func (t *tile) MatchesUp(other Tile) bool {
	otherTile, _ := other.(*tile)

	return t.Up == otherTile.Down
}

func (t *tile) MatchesDown(other Tile) bool {
	otherTile, _ := other.(*tile)

	return t.Down == otherTile.Up
}

func (t *tile) AdjustToRight(other Tile) bool {
	otherTile, _ := other.(*tile)

	if t.MatchesRight(otherTile) {
		// Skip
	} else if t.Down == otherTile.Left {
		t.Rotate90()
	} else if t.Left == otherTile.Left {
		t.Rotate180()
	} else if t.Up == otherTile.Left {
		t.Rotate270()
	} else {
		return false
	}

	return true
}

func (t *tile) AdjustToUp(other Tile) bool {
	otherTile, _ := other.(*tile)

	if t.MatchesUp(otherTile) {
		// Skip
	} else if t.Right == otherTile.Down {
		t.Rotate90()
	} else if t.Down == otherTile.Down {
		t.Rotate180()
	} else if t.Left == otherTile.Down {
		t.Rotate270()
	} else {
		return false
	}

	return true
}

func (t *tile) AdjustToLeft(other Tile) bool {
	otherTile, _ := other.(*tile)

	if t.MatchesLeft(otherTile) {
		// Skip
	} else if t.Up == otherTile.Right {
		t.Rotate90()
	} else if t.Right == otherTile.Right {
		t.Rotate180()
	} else if t.Down == otherTile.Right {
		t.Rotate270()
	} else {
		return false
	}

	return true
}

func (t *tile) AdjustToDown(other Tile) bool {
	otherTile, _ := other.(*tile)

	if t.MatchesDown(otherTile) {
		// Skip
	} else if t.Left == otherTile.Up {
		t.Rotate90()
	} else if t.Up == otherTile.Up {
		t.Rotate180()
	} else if t.Right == otherTile.Up {
		t.Rotate270()
	} else {
		return false
	}

	return true
}
