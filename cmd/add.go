package cmd

import (
	"fmt"

	"github.com/paconte/torneito/model"
)

func AddGame() model.Game {
	country := readCountry()
	name := readName()

	game := model.Game{Country: country, Name: name}
	games.Add(game)

	return game
}

func readName() string {
	var name string
	fmt.Print("Enter name: ")
	fmt.Scanln(&name)
	return name
}

func readCountry() string {
	var country string
	fmt.Print("Enter country: ")
	fmt.Scanln(&country)
	return country
}
