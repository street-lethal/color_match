package main

import (
	"color_match/src"
	"fmt"
)

func main() {
	src.DefaultBoard.Display()

	boards := src.Find(10)
	if boards != nil {
		for _, board := range *boards {
			board.Display()
			fmt.Println("")
		}
	}
}
