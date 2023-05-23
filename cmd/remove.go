package cmd

import "fmt"

func RemoveGame() {
	index := readIndex()
	if index < 0 || index >= len(games) {
		fmt.Println("Invalid index")
	}
	games.Remove(index)
}

func readIndex() int {
	var index int
	fmt.Print("Enter index: ")
	fmt.Scanln(&index)
	return index
}
