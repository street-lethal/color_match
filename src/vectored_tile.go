package src

import "fmt"

type vectoredTile struct {
	Right, Up, Left, Down                 Color
	RightHead, UpHead, LeftHead, DownHead bool
}

func NewVectoredTile(right, up, left, down Color, rightHead, upHead, leftHead, downHead bool) Tile {
	return &vectoredTile{
		right, up, left, down,
		rightHead, upHead, leftHead, downHead,
	}
}

func (t *vectoredTile) Copy() Tile {
	return NewVectoredTile(t.Right, t.Up, t.Left, t.Down, t.RightHead, t.UpHead, t.LeftHead, t.DownHead)
}

func (t *vectoredTile) Display() []string {
	var preRight, preUp, preLeft, preDown string
	if t.RightHead {
		preRight = "#"
	} else {
		preRight = " "
	}
	if t.UpHead {
		preUp = "#"
	} else {
		preUp = " "
	}
	if t.LeftHead {
		preLeft = "#"
	} else {
		preLeft = " "
	}
	if t.DownHead {
		preDown = "#"
	} else {
		preDown = " "
	}
	return []string{
		fmt.Sprintf("  %v%v   ", preUp, t.Up),
		fmt.Sprintf("%v%v  %v%v ", preLeft, t.Left, preRight, t.Right),
		fmt.Sprintf("  %v%v   ", preDown, t.Down),
	}
}

func (t *vectoredTile) OverlapsWith(other Tile) bool {
	otherTile, _ := other.Copy().(*vectoredTile)

	return t.Right == otherTile.Right &&
		t.Down == otherTile.Down &&
		t.Left == otherTile.Left &&
		t.Up == otherTile.Up &&
		t.RightHead == otherTile.RightHead &&
		t.DownHead == otherTile.DownHead &&
		t.LeftHead == otherTile.LeftHead &&
		t.UpHead == otherTile.UpHead
}

func (t *vectoredTile) Rotate90() {
	t.Right, t.Up, t.Left, t.Down = t.Down, t.Right, t.Up, t.Left
	t.RightHead, t.UpHead, t.LeftHead, t.DownHead = t.DownHead, t.RightHead, t.UpHead, t.LeftHead
}

func (t *vectoredTile) Rotate180() {
	t.Right, t.Up, t.Left, t.Down = t.Left, t.Down, t.Right, t.Up
	t.RightHead, t.UpHead, t.LeftHead, t.DownHead = t.LeftHead, t.DownHead, t.RightHead, t.UpHead
}

func (t *vectoredTile) Rotate270() {
	t.Right, t.Up, t.Left, t.Down = t.Up, t.Left, t.Down, t.Right
	t.RightHead, t.UpHead, t.LeftHead, t.DownHead = t.UpHead, t.LeftHead, t.DownHead, t.RightHead
}

func (t *vectoredTile) MatchesRight(other Tile) bool {
	otherTile, _ := other.(*vectoredTile)

	return t.Right == otherTile.Left && t.RightHead == !otherTile.LeftHead
}

func (t *vectoredTile) MatchesLeft(other Tile) bool {
	otherTile, _ := other.(*vectoredTile)

	return t.Left == otherTile.Right && t.LeftHead == !otherTile.RightHead
}

func (t *vectoredTile) MatchesUp(other Tile) bool {
	otherTile, _ := other.(*vectoredTile)

	return t.Up == otherTile.Down && t.UpHead == !otherTile.DownHead
}

func (t *vectoredTile) MatchesDown(other Tile) bool {
	otherTile, _ := other.(*vectoredTile)

	return t.Down == otherTile.Up && t.DownHead == !otherTile.UpHead
}

func (t *vectoredTile) AdjustToRight(other Tile) bool {
	otherTile, _ := other.(*vectoredTile)

	if t.MatchesRight(otherTile) {
		// Skip
	} else if t.Down == otherTile.Left && t.DownHead == !otherTile.LeftHead {
		t.Rotate90()
	} else if t.Left == otherTile.Left && t.LeftHead == !otherTile.LeftHead {
		t.Rotate180()
	} else if t.Up == otherTile.Left && t.UpHead == !otherTile.LeftHead {
		t.Rotate270()
	} else {
		return false
	}

	return true
}

func (t *vectoredTile) AdjustToUp(other Tile) bool {
	otherTile, _ := other.(*vectoredTile)

	if t.MatchesUp(otherTile) {
		// Skip
	} else if t.Right == otherTile.Down && t.RightHead == !otherTile.DownHead {
		t.Rotate90()
	} else if t.Down == otherTile.Down && t.DownHead == !otherTile.DownHead {
		t.Rotate180()
	} else if t.Left == otherTile.Down && t.LeftHead == !otherTile.DownHead {
		t.Rotate270()
	} else {
		return false
	}

	return true
}

func (t *vectoredTile) AdjustToLeft(other Tile) bool {
	otherTile, _ := other.(*vectoredTile)

	if t.MatchesLeft(otherTile) {
		// Skip
	} else if t.Up == otherTile.Right && t.UpHead == !otherTile.RightHead {
		t.Rotate90()
	} else if t.Right == otherTile.Right && t.RightHead == !otherTile.RightHead {
		t.Rotate180()
	} else if t.Down == otherTile.Right && t.DownHead == !otherTile.RightHead {
		t.Rotate270()
	} else {
		return false
	}

	return true
}

func (t *vectoredTile) AdjustToDown(other Tile) bool {
	otherTile, _ := other.(*vectoredTile)

	if t.MatchesDown(otherTile) {
		// Skip
	} else if t.Left == otherTile.Up && t.LeftHead == !otherTile.UpHead {
		t.Rotate90()
	} else if t.Up == otherTile.Up && t.UpHead == !otherTile.UpHead {
		t.Rotate180()
	} else if t.Right == otherTile.Up && t.RightHead == !otherTile.UpHead {
		t.Rotate270()
	} else {
		return false
	}

	return true
}
