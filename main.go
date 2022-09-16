package main

import (
	"color_match/src"
)

func main() {
	src.DefaultBoard.Display()

	boards := src.Find(2)
	if boards != nil {
		for _, board := range *boards {
			board.Display()
		}
	}
}
