package src

import "fmt"

type Board struct {
	UpperLeft, Up, UpperRight,
	Left, Center, Right,
	LowerLeft, Down, LowerRight Tile
}

func (b *Board) Tiles() *[]Tile {
	tiles := []Tile{
		b.UpperLeft, b.Up, b.UpperRight,
		b.Left, b.Center, b.Right,
		b.LowerLeft, b.Down, b.LowerRight,
	}

	return &tiles
}

func (b *Board) Display() {
	rows := [][]Tile{
		{b.UpperLeft, b.Up, b.UpperRight},
		{b.Left, b.Center, b.Right},
		{b.LowerLeft, b.Down, b.LowerRight},
	}
	fmt.Println("+-------+-------+-------+")
	for _, tiles := range rows {
		left, center, right := tiles[0], tiles[1], tiles[2]
		fmt.Printf("|%v|%v|%v|\n", left.Display()[0], center.Display()[0], right.Display()[0])
		fmt.Printf("|%v|%v|%v|\n", left.Display()[1], center.Display()[1], right.Display()[1])
		fmt.Printf("|%v|%v|%v|\n", left.Display()[2], center.Display()[2], right.Display()[2])
		fmt.Println("+-------+-------+-------+")
	}
}

func (b *Board) ReOrder(ids []int) {
	tiles := *b.Tiles()
	b.UpperLeft, b.Up, b.UpperRight,
		b.Left, b.Center, b.Right,
		b.LowerLeft, b.Down, b.LowerRight =
		tiles[ids[0]], tiles[ids[1]], tiles[ids[2]],
		tiles[ids[3]], tiles[ids[4]], tiles[ids[5]],
		tiles[ids[6]], tiles[ids[7]], tiles[ids[8]]
}

func (b *Board) Copy() *Board {
	return &Board{
		b.UpperLeft.Copy(), b.Up.Copy(), b.UpperRight.Copy(),
		b.Left.Copy(), b.Center.Copy(), b.Right.Copy(),
		b.LowerLeft.Copy(), b.Down.Copy(), b.LowerRight.Copy(),
	}
}

func (b *Board) IsEquivalentTo(other *Board) bool {
	copied := other.Copy()

	for i := 0; i < 4; i++ {
		if b.Center.OverlapsWith(copied.Center) {
			return b.UpperLeft.OverlapsWith(copied.UpperLeft) &&
				b.Up.OverlapsWith(copied.Up) &&
				b.UpperRight.OverlapsWith(copied.UpperRight) &&
				b.Left.OverlapsWith(copied.Left) &&
				b.Center.OverlapsWith(copied.Center) &&
				b.Right.OverlapsWith(copied.Right) &&
				b.LowerLeft.OverlapsWith(copied.LowerLeft) &&
				b.Down.OverlapsWith(copied.Down) &&
				b.LowerRight.OverlapsWith(copied.LowerRight)
		}
		copied.rotate()
	}

	return false
}

func (b *Board) rotate() {
	b.UpperLeft, b.Up, b.UpperRight,
		b.Right, b.LowerRight,
		b.Down, b.LowerLeft, b.Left =
		b.UpperRight, b.Right, b.LowerRight,
		b.Down, b.LowerLeft,
		b.Left, b.UpperLeft, b.Up
	b.UpperLeft.Rotate90()
	b.Up.Rotate90()
	b.UpperRight.Rotate90()
	b.Left.Rotate90()
	b.Center.Rotate90()
	b.Right.Rotate90()
	b.LowerLeft.Rotate90()
	b.Down.Rotate90()
	b.LowerRight.Rotate90()
}

func (b *Board) adjustAround() bool {
	return b.Right.AdjustToLeft(b.Center) &&
		b.Up.AdjustToDown(b.Center) &&
		b.Left.AdjustToRight(b.Center) &&
		b.Down.AdjustToUp(b.Center)
}

func (b *Board) AdjustCenterAndAround() bool {
	for i := 0; i < 4; i++ {
		if b.adjustAround() {
			return true
		}

		b.Center.Rotate90()
	}
	return false
}

func (b *Board) adjustUpperLeft() bool {
	return b.UpperLeft.AdjustToRight(b.Up) &&
		b.UpperLeft.MatchesDown(b.Left)
}

func (b *Board) adjustUpperRight() bool {
	return b.UpperRight.AdjustToLeft(b.Up) &&
		b.UpperRight.MatchesDown(b.Right)
}

func (b *Board) adjustLowerLeft() bool {
	return b.LowerLeft.AdjustToRight(b.Down) &&
		b.LowerLeft.MatchesUp(b.Left)
}

func (b *Board) adjustLowerRight() bool {
	return b.LowerRight.AdjustToLeft(b.Down) &&
		b.LowerRight.MatchesUp(b.Right)
}

func (b *Board) AdjustCorners() bool {
	return b.adjustUpperLeft() &&
		b.adjustUpperRight() &&
		b.adjustLowerLeft() &&
		b.adjustLowerRight()
}

func (b *Board) Adjust() bool {
	return b.AdjustCenterAndAround() && b.AdjustCorners()
}

func (b *Board) IsCorrect() bool {
	return b.Center.MatchesUp(b.Up) &&
		b.Center.MatchesLeft(b.Left) &&
		b.Center.MatchesDown(b.Down) &&
		b.Center.MatchesRight(b.Right) &&
		b.UpperLeft.MatchesRight(b.Up) &&
		b.UpperLeft.MatchesDown(b.Left) &&
		b.UpperRight.MatchesLeft(b.Up) &&
		b.UpperRight.MatchesDown(b.Right) &&
		b.LowerLeft.MatchesUp(b.Left) &&
		b.LowerLeft.MatchesRight(b.Down) &&
		b.LowerRight.MatchesUp(b.Right) &&
		b.LowerRight.MatchesLeft(b.Down)
}
