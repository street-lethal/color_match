package src

type Color string

const (
	Red    Color = "R"
	Yellow Color = "Y"
	Green  Color = "G"
	Blue   Color = "B"
	Pink   Color = "P"
)

type Tile struct {
	Right, Up, Left, Down Color
}

func NewTile(right, up, left, down Color) *Tile {
	return &Tile{right, up, left, down}
}

func (t *Tile) Copy() *Tile {
	return NewTile(t.Right, t.Up, t.Left, t.Down)
}

func (t *Tile) Rotate90() {
	t.Right, t.Up, t.Left, t.Down = t.Down, t.Right, t.Up, t.Left
}

func (t *Tile) Rotate180() {
	t.Right, t.Up, t.Left, t.Down = t.Left, t.Down, t.Right, t.Up
}

func (t *Tile) Rotate270() {
	t.Right, t.Up, t.Left, t.Down = t.Up, t.Left, t.Down, t.Right
}

func (t *Tile) MatchesRight(other *Tile) bool {
	return t.Right == other.Left
}

func (t *Tile) MatchesLeft(other *Tile) bool {
	return t.Left == other.Right
}

func (t *Tile) MatchesUp(other *Tile) bool {
	return t.Up == other.Down
}

func (t *Tile) MatchesDown(other *Tile) bool {
	return t.Down == other.Up
}

func (t *Tile) AdjustToRight(other *Tile) bool {
	if t.MatchesRight(other) {
		// Skip
	} else if t.Down == other.Left {
		t.Rotate90()
	} else if t.Left == other.Left {
		t.Rotate180()
	} else if t.Up == other.Left {
		t.Rotate270()
	} else {
		return false
	}

	return true
}

func (t *Tile) AdjustToUp(other *Tile) bool {
	if t.MatchesUp(other) {
		// Skip
	} else if t.Right == other.Down {
		t.Rotate90()
	} else if t.Down == other.Down {
		t.Rotate180()
	} else if t.Left == other.Down {
		t.Rotate270()
	} else {
		return false
	}

	return true
}

func (t *Tile) AdjustToLeft(other *Tile) bool {
	if t.MatchesLeft(other) {
		// Skip
	} else if t.Up == other.Right {
		t.Rotate90()
	} else if t.Right == other.Right {
		t.Rotate180()
	} else if t.Down == other.Right {
		t.Rotate270()
	} else {
		return false
	}

	return true
}

func (t *Tile) AdjustToDown(other *Tile) bool {
	if t.MatchesDown(other) {
		// Skip
	} else if t.Left == other.Up {
		t.Rotate90()
	} else if t.Up == other.Up {
		t.Rotate180()
	} else if t.Right == other.Up {
		t.Rotate270()
	} else {
		return false
	}

	return true
}
