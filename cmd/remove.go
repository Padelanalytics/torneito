package cmd

import (
	"fmt"
	"slices"
)

func RemoveGame() {
	index, err := readIndex()
	if err != nil {
		fmt.Println(err)
		return
	}
	games = slices.Delete(games, index, index)
}

func readIndex() (int, error) {
	var index int
	fmt.Print("Enter game index: ")
	fmt.Scanln(&index)
	if index < 0 || index >= len(games) {
		return -1, fmt.Errorf("invalid game selection")
	}
	return index, nil
}
