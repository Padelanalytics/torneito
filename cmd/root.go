package cmd

import (
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
	"github.com/paconte/torneito/model"
)

var games model.Games = model.Games{}

func Execute() {
	for {
		selectAction()
	}
}

func selectAction() {
	prompt := promptui.Select{
		Label: "Select action",
		Items: []string{"Add", "Delete", "List", "Export", "Exit"},
	}
	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	switch result {
	case "Add":
		AddGame()
	case "Delete":
		RemoveGame()
	case "List":
		ListGames()
	case "Export":
		AddGame()
	case "Exit":
		fmt.Println("Bye!")
		os.Exit(1)
	}
}
