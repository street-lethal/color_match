package src

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

func Find(count int) *[]Board {
	ids := []int{8, 7, 6, 5, 4, 3, 2, 1, 0}

	correctBoards := make([]Board, 0)

	for ids != nil {
		board := &DefaultBoard
		board.ReOrder(ids)
		if board.Adjust() {
			alreadyFound := false
			for _, foundBoard := range correctBoards {
				if board.IsEquivalentTo(&foundBoard) {
					alreadyFound = true
					break
				}
			}

			if alreadyFound {
				continue
			}

			correctBoards = append(correctBoards, *board.Copy())
			if len(correctBoards) >= count {
				break
			}
		}
		ids = getNext(ids)
	}

	return &correctBoards
}
