package cmd

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

func EditGame() {
	index, err := readIndex()
	if err != nil {
		fmt.Println(err)
		return
	}
	selectEditAction(index)
}

func selectEditAction(index int) {
	prompt := promptui.Select{
		Label: "Select action",
		Items: []string{"Federation", "Name", "Serie", "Division", "Date", "Round", "Category", "P1 Last", "P1 First", "P3 Last", "P4 First", "Scores"},
	}
	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	game := games[index]

	switch result {
	case "Federation":
		input := read(tournaments.Countries(), "Enter country:")
		game.Country = input
	case "Name":
		input := read(tournaments.Names(), "Enter name:")
		game.Name = input
	case "Serie":
		input := read(tournaments.Series(), "Enter serie:")
		game.Serie = input
	case "Division":
		input := read(tournaments.Divisions(), "Enter division:")
		game.Division = input
	case "Date":
		input := read(games.Dates(), "Enter date:")
		game.Date = input
	case "Round":
		input := read(games.Rounds(), "Enter round:")
		game.Round = input
	case "Category":
		input := read(games.Categories(), "Enter category:")
		game.Category = input
	case "P1 Last":
		input := read(games.Lastnames(), "Enter player 1 last name:")
		game.Players[0] = input
	case "P1 First":
		input := read(games.Firstnames(), "Enter player 1 first name:")
		game.Players[1] = input
	case "P2 Last":
		input := read(games.Lastnames(), "Enter player 2 last name:")
		game.Players[2] = input
	case "P2 First":
		input := read(games.Firstnames(), "Enter player 2 first name:")
		game.Players[3] = input
	case "P3 Last":
		input := read(games.Lastnames(), "Enter player 3 last name:")
		game.Players[4] = input
	case "P3 First":
		input := read(games.Firstnames(), "Enter player 3 first name:")
		game.Players[5] = input
	case "P4 Last":
		input := read(games.Lastnames(), "Enter player 4 last name:")
		game.Players[6] = input
	case "P4 First":
		input := read(games.Firstnames(), "Enter player 4 first name:")
		game.Players[7] = input
	case "Scores":
		input := readScores()
		game.Scores = input
	case "":
		fmt.Println("Skipping!")
		return
	}

	games.Update(index, game)
}
