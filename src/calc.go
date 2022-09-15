package src

func Solve() {
	b := DefaultBoard
	b.Display()

	board := find()
	if board != nil {
		board.Display()
	}
}

func getNext(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		if nums[i-1] < nums[i] {
			continue
		}

		sorted := make([]int, i)
		copy(sorted, nums[0:i])

		j := i - 1
		m := nums[j]
		for tmp := i - 2; tmp >= 0; tmp-- {
			if nums[tmp] < nums[i] {
				break
			}

			j = tmp
			m = nums[tmp]
		}

		sorted[j] = nums[i]
		for k := 0; k < i/2; k++ {
			x := i - k - 1
			sorted[k], sorted[x] = sorted[x], sorted[k]
		}

		ret := append([]int{}, sorted...)
		ret = append(ret, m)
		ret = append(ret, nums[i+1:]...)
		return ret
	}

	return nil
}

func find() *Board {
	var board *Board

	ids := []int{8, 7, 6, 5, 4, 3, 2, 1, 0}

	for ids != nil {
		tiles := make([]*Tile, 9)
		for i, id := range ids {
			tile := (*DefaultBoard.Tiles())[id]
			tiles[i] = tile.Copy()
		}
		board = &Board{
			tiles[0],
			tiles[1],
			tiles[2],
			tiles[3],
			tiles[4],
			tiles[5],
			tiles[6],
			tiles[7],
			tiles[8],
		}
		if board.Adjust() {
			return board
		}
		ids = getNext(ids)
	}

	return nil
}
