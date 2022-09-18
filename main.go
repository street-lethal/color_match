package main

import (
	"color_match/src"
	"fmt"
)

func main() {
	src.DefaultVectoredBoard.Display()

	boards := src.Find(2)
	if boards != nil {
		for _, board := range *boards {
			board.Display()
			fmt.Println("")
		}
	}
}
